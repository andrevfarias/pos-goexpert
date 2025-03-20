package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/middleware/ratelimiter"
)

type RedisApiKeyCache struct {
	redisClient *redis.Client
}

func NewRedisApiKeyCache(redisClient *redis.Client) *RedisApiKeyCache {
	return &RedisApiKeyCache{redisClient: redisClient}
}

func (r *RedisApiKeyCache) GetApiKey(apiKey string) (ratelimiter.ApiKey, error) {
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

func (r *RedisApiKeyCache) InsertOrUpdateApiKey(apiKey ratelimiter.ApiKey) error {
	key := fmt.Sprintf("apiKey:%s", apiKey.Key)
	return r.redisClient.Set(context.Background(), key, apiKey.RateLimit, 0).Err()
}

func (r *RedisApiKeyCache) DeleteApiKey(apiKey string) error {
	key := fmt.Sprintf("apiKey:%s", apiKey)
	return r.redisClient.Del(context.Background(), key).Err()
}
