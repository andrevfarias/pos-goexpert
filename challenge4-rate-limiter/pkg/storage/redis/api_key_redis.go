package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
)

type RedisApiKeyStorage struct {
	redisClient *redis.Client
}

func NewRedisApiKeyStorage(redisClient *redis.Client) *RedisApiKeyStorage {
	return &RedisApiKeyStorage{redisClient: redisClient}
}

func (r *RedisApiKeyStorage) GetApiKey(apiKey string) (ratelimiter.ApiKey, error) {
	key := fmt.Sprintf("apiKey:%s", apiKey)
	rateLimit, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return ratelimiter.ApiKey{}, ratelimiter.ErrAPIKeyNotFound
		}
		return ratelimiter.ApiKey{}, err
	}

	rateLimitInt, err := strconv.Atoi(rateLimit)
	if err != nil {
		return ratelimiter.ApiKey{}, err
	}

	return ratelimiter.ApiKey{
		Key:       apiKey,
		RateLimit: rateLimitInt,
	}, nil
}

func (r *RedisApiKeyStorage) InsertOrUpdateApiKey(apiKey ratelimiter.ApiKey) error {
	key := fmt.Sprintf("apiKey:%s", apiKey.Key)
	return r.redisClient.Set(context.Background(), key, apiKey.RateLimit, 0).Err()
}

func (r *RedisApiKeyStorage) DeleteApiKey(apiKey string) error {
	key := fmt.Sprintf("apiKey:%s", apiKey)
	return r.redisClient.Del(context.Background(), key).Err()
}
