package memory

import (
	"fmt"
	"sync"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
)

type MemoryClientStateStorage struct {
	clients map[string]ratelimiter.ClientState
	mu      sync.RWMutex
}

func NewMemoryClientStateStorage() *MemoryClientStateStorage {
	return &MemoryClientStateStorage{
		clients: make(map[string]ratelimiter.ClientState),
	}
}

func (r *MemoryClientStateStorage) GetClientState(clientID string, clientType ratelimiter.ClientType) (ratelimiter.ClientState, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	clientState, ok := r.clients[key]
	if !ok {
		return ratelimiter.ClientState{}, ratelimiter.ErrClientStateNotFound
	}

	return clientState, nil
}

func (r *MemoryClientStateStorage) InsertOrUpdateClientState(clientID string, clientState ratelimiter.ClientState, clientType ratelimiter.ClientType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	clientState.WindowStart = clientState.WindowStart.UTC()
	clientState.BlockUntil = clientState.BlockUntil.UTC()

	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	r.clients[key] = clientState
	return nil
}

func (r *MemoryClientStateStorage) DeleteClientState(clientID string, clientType ratelimiter.ClientType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	delete(r.clients, key)
	return nil
}
