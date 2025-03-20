package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/middleware/ratelimiter"
	"github.com/redis/go-redis/v9"
)

type RedisClientStateCache struct {
	redisClient *redis.Client
}

func NewRedisClientStateCache(redisClient *redis.Client) *RedisClientStateCache {
	return &RedisClientStateCache{redisClient: redisClient}
}

func (r *RedisClientStateCache) GetClientState(clientID string, clientType ratelimiter.ClientType) (ratelimiter.ClientState, error) {
	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	clientStateJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return ratelimiter.ClientState{}, ratelimiter.ErrClientStateNotFound
	}
	if err != nil {
		return ratelimiter.ClientState{}, err
	}

	var clientState ratelimiter.ClientState
	err = json.Unmarshal([]byte(clientStateJson), &clientState)
	if err != nil {
		return ratelimiter.ClientState{}, err
	}

	return clientState, nil
}

func (r *RedisClientStateCache) InsertOrUpdateClientState(clientID string, clientState ratelimiter.ClientState, clientType ratelimiter.ClientType) error {
	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	clientStateJson, err := json.Marshal(clientState)
	if err != nil {
		return err
	}

	return r.redisClient.Set(context.Background(), key, clientStateJson, 0).Err()
}

func (r *RedisClientStateCache) DeleteClientState(clientID string, clientType ratelimiter.ClientType) error {
	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	return r.redisClient.Del(context.Background(), key).Err()
}
