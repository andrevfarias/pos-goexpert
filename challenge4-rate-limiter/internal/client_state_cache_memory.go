package internal

import (
	"fmt"
	"sync"

	ratelimiter "github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/middleware/rate-limiter"
)

type InMemoryClientStateCache struct {
	clients map[string]*ratelimiter.ClientState
	mu      sync.RWMutex
}

func NewInMemoryClientStateCache() *InMemoryClientStateCache {
	return &InMemoryClientStateCache{
		clients: make(map[string]*ratelimiter.ClientState),
	}
}

func (r *InMemoryClientStateCache) GetClientState(clientID string, clientType ratelimiter.ClientType) (*ratelimiter.ClientState, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := fmt.Sprintf("%s:%s", clientType, clientID)
	clientState, ok := r.clients[key]
	if !ok {
		return nil, ratelimiter.ErrClientStateNotFound
	}
	clientStateCopy := *clientState
	return &clientStateCopy, nil
}

func (r *InMemoryClientStateCache) InsertOrUpdateClientState(clientID string, clientState ratelimiter.ClientState, clientType ratelimiter.ClientType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := fmt.Sprintf("%s:%s", clientType, clientID)
	clientStateCopy := clientState
	r.clients[key] = &clientStateCopy
	return nil
}

func (r *InMemoryClientStateCache) DeleteClientState(clientID string, clientType ratelimiter.ClientType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := fmt.Sprintf("%s:%s", clientType, clientID)
	delete(r.clients, key)
	return nil
}
