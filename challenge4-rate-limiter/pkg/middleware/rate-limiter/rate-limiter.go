package ratelimiter

import (
	"errors"
	"net"
	"net/http"
	"time"
)

type rateLimiter struct {
	IPRateLimit      int
	BlockTime        time.Duration
	ClientStateCache ClientStateCache
	ApiKeyCache      ApiKeyCache
}

func New(config RateLimiterConfig) func(next http.Handler) http.Handler {
	rl := &rateLimiter{
		IPRateLimit:      config.IPRateLimit,
		BlockTime:        config.BlockTime,
		ClientStateCache: config.ClientStateCache,
		ApiKeyCache:      config.ApiKeyCache,
	}

	return rl.Handler
}

func (rl *rateLimiter) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiKey := r.Header.Get("API_KEY"); apiKey != "" {
			isAllowed, err := rl.isApiKeyAllowed(apiKey)

			if err == nil && !isAllowed {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			if err == nil && isAllowed {
				next.ServeHTTP(w, r)
				return
			}

			if err != nil && err != ErrAPIKeyNotFound {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if (err != nil) || (ip == "") {
			http.Error(w, "cannot resolve origin", http.StatusInternalServerError)
			return
		}

		isAllowed, err := rl.isIpAllowed(ip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !isAllowed {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *rateLimiter) isIpAllowed(ip string) (bool, error) {
	if ip == "" {
		return false, errors.New("cannot resolve origin")
	}

	ipClient, err := rl.ClientStateCache.GetClientState(ip, ClientTypes.IP)
	if err != nil {
		return false, err
	}

	if ipClient == nil {
		ipClient = &ClientState{
			WindowStart:  time.Now(),
			RequestCount: 0,
			Blocked:      false,
		}
	}

	isAllowed, err := rl.isAllowed(rl.IPRateLimit, ipClient)
	if err != nil {
		return false, err
	}

	rl.ClientStateCache.InsertOrUpdateClientState(ip, *ipClient, ClientTypes.IP)

	return isAllowed, nil
}

func (rl *rateLimiter) isApiKeyAllowed(apiKey string) (bool, error) {
	token, err := rl.ApiKeyCache.GetApiKey(apiKey)
	if err != nil {
		return false, err
	}
	if token.Key == "" {
		return false, ErrAPIKeyNotFound
	}

	clientState, err := rl.ClientStateCache.GetClientState(apiKey, ClientTypes.ApiKey)
	if err != nil {
		return false, err
	}

	if clientState == nil {
		clientState = &ClientState{
			WindowStart:  time.Now(),
			RequestCount: 0,
			Blocked:      false,
		}
	}
	isAllowed, err := rl.isAllowed(token.RateLimit, clientState)
	if err != nil {
		return false, err
	}

	rl.ClientStateCache.InsertOrUpdateClientState(apiKey, *clientState, ClientTypes.ApiKey)

	return isAllowed, nil
}

func (rl *rateLimiter) isAllowed(rateLimit int, client *ClientState) (bool, error) {
	now := time.Now()

	if client.Blocked {
		if now.Before(client.BlockUntil) {
			return false, nil
		}
		client.Blocked = false
		client.WindowStart = now
		client.RequestCount = 0
	}

	timeWindowSub := now.Sub(client.WindowStart)
	if timeWindowSub > time.Second {
		client.WindowStart = now
		client.RequestCount = 0
	}

	client.RequestCount++
	if client.RequestCount > rateLimit {
		client.Blocked = true
		client.BlockUntil = now.Add(rl.BlockTime)
		return false, nil
	}

	return true, nil
}
