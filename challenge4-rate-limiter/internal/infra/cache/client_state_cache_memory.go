package cache

import (
	"fmt"
	"sync"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/middleware/ratelimiter"
)

type InMemoryClientStateCache struct {
	clients map[string]ratelimiter.ClientState
	mu      sync.RWMutex
}

func NewInMemoryClientStateCache() *InMemoryClientStateCache {
	return &InMemoryClientStateCache{
		clients: make(map[string]ratelimiter.ClientState),
	}
}

func (r *InMemoryClientStateCache) GetClientState(clientID string, clientType ratelimiter.ClientType) (ratelimiter.ClientState, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := fmt.Sprintf("%s:%s", clientType, clientID)
	clientState, ok := r.clients[key]
	if !ok {
		return ratelimiter.ClientState{}, ratelimiter.ErrClientStateNotFound
	}

	return clientState, nil
}

func (r *InMemoryClientStateCache) InsertOrUpdateClientState(clientID string, clientState ratelimiter.ClientState, clientType ratelimiter.ClientType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := fmt.Sprintf("%s:%s", clientType, clientID)
	r.clients[key] = clientState
	return nil
}

func (r *InMemoryClientStateCache) DeleteClientState(clientID string, clientType ratelimiter.ClientType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := fmt.Sprintf("%s:%s", clientType, clientID)
	delete(r.clients, key)
	return nil
}
