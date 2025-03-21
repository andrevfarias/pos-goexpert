package redis

import (
	"context"
	"testing"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type ApiKeyRedisStorageTestSuite struct {
	suite.Suite
	redisClient *redis.Client
	storage     *RedisApiKeyStorage
}

func (s *ApiKeyRedisStorageTestSuite) SetupTest() {

	s.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()
	if err := s.redisClient.Ping(ctx).Err(); err != nil {
		s.T().Fatalf("failed to connect to Redis: %s", err)
	}
	s.redisClient.FlushAll(ctx)

	s.storage = NewRedisApiKeyStorage(s.redisClient)
}

func (s *ApiKeyRedisStorageTestSuite) TearDownTest() {
	ctx := context.Background()
	s.redisClient.FlushAll(ctx)
	s.redisClient.Close()
}

func TestApiKeyRedisStorage(t *testing.T) {
	suite.Run(t, new(ApiKeyRedisStorageTestSuite))
}

func (s *ApiKeyRedisStorageTestSuite) TestShouldReturnErrorWhenApiKeyIsNotFound() {
	_, err := s.storage.GetApiKey("test")
	s.Error(err)
}

func (s *ApiKeyRedisStorageTestSuite) TestShouldReturnApiKeyWhenItExists() {
	s.storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 10})
	apiKey, err := s.storage.GetApiKey("test")
	s.NoError(err)
	s.Equal("test", apiKey.Key)
	s.Equal(10, apiKey.RateLimit)
}

func (s *ApiKeyRedisStorageTestSuite) TestShouldUpdateApiKeyWhenItExists() {
	s.storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 10})
	s.storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 20})
	apiKey, err := s.storage.GetApiKey("test")
	s.NoError(err)
	s.Equal("test", apiKey.Key)
	s.Equal(20, apiKey.RateLimit)
}

func (s *ApiKeyRedisStorageTestSuite) TestShouldDeleteApiKeyWhenItExists() {
	s.storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 10})
	s.storage.DeleteApiKey("test")
	_, err := s.storage.GetApiKey("test")
	s.Error(err)
}
