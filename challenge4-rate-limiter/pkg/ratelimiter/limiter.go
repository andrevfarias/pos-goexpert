package ratelimiter

import (
	"errors"
	"time"
)

type rateLimiter struct {
	IPRateLimit      int
	BlockTime        time.Duration
	ClientStateCache ClientStateCache
	ApiKeyCache      ApiKeyCache
}

func NewRateLimiter(config RateLimiterConfig) RateLimiterService {
	return &rateLimiter{
		IPRateLimit:      config.IPRateLimit,
		BlockTime:        config.BlockTime,
		ClientStateCache: config.ClientStateCache,
		ApiKeyCache:      config.ApiKeyCache,
	}
}

func (rl *rateLimiter) IsIpAllowed(ip string) (bool, error) {
	if ip == "" {
		return false, errors.New("cannot resolve origin")
	}

	ipClient, err := rl.ClientStateCache.GetClientState(ip, ClientTypes.IP)
	if err != nil {
		if err != ErrClientStateNotFound {
			return false, err
		}
		ipClient = ClientState{
			WindowStart:  time.Now(),
			RequestCount: 0,
			Blocked:      false,
		}
	}

	isAllowed, err := rl.isAllowed(rl.IPRateLimit, &ipClient)
	if err != nil {
		return false, err
	}

	rl.ClientStateCache.InsertOrUpdateClientState(ip, ipClient, ClientTypes.IP)

	return isAllowed, nil
}

func (rl *rateLimiter) IsApiKeyAllowed(apiKey string) (bool, error) {
	token, err := rl.ApiKeyCache.GetApiKey(apiKey)
	if err != nil {
		return false, err
	}

	if token.Key == "" {
		return false, ErrAPIKeyNotFound
	}

	clientState, err := rl.ClientStateCache.GetClientState(apiKey, ClientTypes.ApiKey)
	if err != nil {
		if err != ErrClientStateNotFound {
			return false, err
		}
		clientState = ClientState{
			WindowStart:  time.Now(),
			RequestCount: 0,
			Blocked:      false,
		}
	}

	isAllowed, err := rl.isAllowed(token.RateLimit, &clientState)
	if err != nil {
		return false, err
	}

	rl.ClientStateCache.InsertOrUpdateClientState(apiKey, clientState, ClientTypes.ApiKey)

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
