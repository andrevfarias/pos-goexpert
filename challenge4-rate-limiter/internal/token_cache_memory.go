package internal

import (
	"sync"

	ratelimiter "github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/middleware/rate-limiter"
)

type InMemoryApiKeyCache struct {
	apiKeys map[string]*ratelimiter.ApiKey
	mu      sync.RWMutex
}

func NewInMemoryApiKeyCache() *InMemoryApiKeyCache {
	return &InMemoryApiKeyCache{
		apiKeys: make(map[string]*ratelimiter.ApiKey),
	}
}

func (r *InMemoryApiKeyCache) GetApiKey(key string) (ratelimiter.ApiKey, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	apiKey, ok := r.apiKeys[key]
	if !ok {
		return ratelimiter.ApiKey{}, ratelimiter.ErrAPIKeyNotFound
	}
	apiKeyCopy := *apiKey
	return apiKeyCopy, nil
}

func (r *InMemoryApiKeyCache) InsertOrUpdateApiKey(apiKey *ratelimiter.ApiKey) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	apiKeyCopy := *apiKey
	r.apiKeys[apiKey.Key] = &apiKeyCopy
	return nil
}

func (r *InMemoryApiKeyCache) DeleteApiKey(apiKey string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.apiKeys, apiKey)
	return nil
}
