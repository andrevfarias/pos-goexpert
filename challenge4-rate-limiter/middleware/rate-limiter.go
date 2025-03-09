package middleware

import (
	"errors"
	"net"
	"net/http"
	"time"
)

var (
	ErrAPIKeyNotFound = errors.New("API_KEY not found")
)

type rateLimiter struct {
	IPRateLimit     int
	APIKeyRateLimit map[string]int
	BlockTime       time.Duration
	IPClients       map[string]*clientState
	APIKeyClients   map[string]*clientState
}

type RateLimiterConfig struct {
	IPRateLimit     int
	APIKeyRateLimit map[string]int
	BlockTime       time.Duration
}

type clientState struct {
	windowStart  time.Time
	requestCount int
	blocked      bool
	blockUntil   time.Time
}

func RateLimiter(config RateLimiterConfig) func(next http.Handler) http.Handler {
	rl := &rateLimiter{
		IPRateLimit:     config.IPRateLimit,
		APIKeyRateLimit: config.APIKeyRateLimit,
		BlockTime:       config.BlockTime,
		IPClients:       make(map[string]*clientState),
		APIKeyClients:   make(map[string]*clientState),
	}

	return rl.Handler
}

func (rl *rateLimiter) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiKey := r.Header.Get("API_KEY"); apiKey != "" {
			isAllowed, err := rl.isApiKeyAllowed(apiKey)
			if err != nil {
				if err == ErrAPIKeyNotFound {
					http.Error(w, "API_KEY not found", http.StatusNotFound)
					return
				}
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if !isAllowed {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if (err != nil) || (ip == "") {
			http.Error(w, "cannot resolve origin", http.StatusInternalServerError)
			return
		}

		isAllowed, err := rl.isIpAllowed(ip)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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

	ipClient, exists := rl.IPClients[ip]
	if !exists {
		ipClient = &clientState{
			windowStart:  time.Now(),
			requestCount: 0,
			blocked:      false,
		}
		rl.IPClients[ip] = ipClient
	}

	return rl.isAllowed(rl.IPRateLimit, ipClient)
}

func (rl *rateLimiter) isApiKeyAllowed(apiKey string) (bool, error) {
	rateLimit, ok := rl.APIKeyRateLimit[apiKey]
	if !ok {
		return false, ErrAPIKeyNotFound
	}

	apiKeyClient, exists := rl.APIKeyClients[apiKey]
	if !exists {
		apiKeyClient = &clientState{
			windowStart:  time.Now(),
			requestCount: 0,
			blocked:      false,
		}
		rl.APIKeyClients[apiKey] = apiKeyClient
	}

	return rl.isAllowed(rateLimit, apiKeyClient)
}

func (rl *rateLimiter) isAllowed(rateLimit int, client *clientState) (bool, error) {
	now := time.Now()

	if client.blocked {
		if now.Before(client.blockUntil) {
			return false, nil
		}
		client.blocked = false
		client.windowStart = now
		client.requestCount = 0
	}

	if now.Sub(client.windowStart) > time.Second {
		client.windowStart = now
		client.requestCount = 0
	}

	client.requestCount++
	if client.requestCount > rateLimit {
		client.blocked = true
		client.blockUntil = now.Add(rl.BlockTime)
		return false, nil
	}

	return true, nil

}
