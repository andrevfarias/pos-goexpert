package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
	"github.com/redis/go-redis/v9"
)

type RedisClientStateStorage struct {
	redisClient *redis.Client
}

func NewRedisClientStateStorage(redisClient *redis.Client) *RedisClientStateStorage {
	return &RedisClientStateStorage{redisClient: redisClient}
}

func (r *RedisClientStateStorage) GetClientState(clientID string, clientType ratelimiter.ClientType) (ratelimiter.ClientState, error) {
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

func (r *RedisClientStateStorage) InsertOrUpdateClientState(clientID string, clientState ratelimiter.ClientState, clientType ratelimiter.ClientType) error {
	clientState.WindowStart = clientState.WindowStart.UTC()
	clientState.BlockUntil = clientState.BlockUntil.UTC()
	clientStateJson, err := json.Marshal(clientState)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	return r.redisClient.Set(context.Background(), key, clientStateJson, 0).Err()
}

func (r *RedisClientStateStorage) DeleteClientState(clientID string, clientType ratelimiter.ClientType) error {
	key := fmt.Sprintf("state:%s:%s", clientType, clientID)
	return r.redisClient.Del(context.Background(), key).Err()
}
