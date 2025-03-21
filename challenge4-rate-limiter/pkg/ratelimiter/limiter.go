package ratelimiter

import (
	"errors"
	"net"
	"time"
)

type rateLimiter struct {
	IPRateLimit        int
	BlockTime          time.Duration
	ClientStateStorage ClientStateStorage
	ApiKeyStorage      ApiKeyStorage
}

func NewRateLimiter(config RateLimiterConfig) RateLimiterService {
	return &rateLimiter{
		IPRateLimit:        config.IPRateLimit,
		BlockTime:          config.BlockTime,
		ClientStateStorage: config.ClientStateStorage,
		ApiKeyStorage:      config.ApiKeyStorage,
	}
}

func (rl *rateLimiter) IsIpAllowed(ip string) (bool, error) {
	if ip == "" || net.ParseIP(ip) == nil {
		return false, errors.New("cannot resolve origin")
	}

	ipClient, err := rl.ClientStateStorage.GetClientState(ip, ClientTypes.IP)
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

	rl.ClientStateStorage.InsertOrUpdateClientState(ip, ipClient, ClientTypes.IP)

	return isAllowed, nil
}

func (rl *rateLimiter) IsApiKeyAllowed(apiKey string) (bool, error) {
	token, err := rl.ApiKeyStorage.GetApiKey(apiKey)
	if err != nil {
		return false, err
	}

	clientState, err := rl.ClientStateStorage.GetClientState(apiKey, ClientTypes.ApiKey)
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

	rl.ClientStateStorage.InsertOrUpdateClientState(apiKey, clientState, ClientTypes.ApiKey)

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
