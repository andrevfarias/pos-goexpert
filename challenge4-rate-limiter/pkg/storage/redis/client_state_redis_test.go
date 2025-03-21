package redis

import (
	"context"
	"testing"
	"time"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type ClientStateRedisTestSuite struct {
	suite.Suite
	redisClient *redis.Client
	storage     *RedisClientStateStorage
	clientType  ratelimiter.ClientType
	key         string
}

func (s *ClientStateRedisTestSuite) SetupTest() {
	s.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()
	if err := s.redisClient.Ping(ctx).Err(); err != nil {
		s.T().Fatalf("failed to connect to Redis: %s", err)
	}
	s.redisClient.FlushAll(ctx)

	s.storage = NewRedisClientStateStorage(s.redisClient)
}

func (s *ClientStateRedisTestSuite) TearDownTest() {
	ctx := context.Background()
	s.redisClient.FlushAll(ctx)
	s.redisClient.Close()
}

func TestClientStateRedisStorage(t *testing.T) {
	testSuite := new(ClientStateRedisTestSuite)

	testSuite.clientType = ratelimiter.ClientTypes.IP
	testSuite.key = "127.0.0.1"
	suite.Run(t, testSuite)

	testSuite.clientType = ratelimiter.ClientTypes.ApiKey
	testSuite.key = "test"
	suite.Run(t, testSuite)
}

func (s *ClientStateRedisTestSuite) TestClientStateStorage() {
	now := time.Now().UTC()

	// Check Insert
	clientState := ratelimiter.ClientState{
		WindowStart:  now,
		RequestCount: 10,
		Blocked:      false,
		BlockUntil:   now,
	}

	err := s.storage.InsertOrUpdateClientState(s.key, clientState, s.clientType)
	s.NoError(err)

	// Check Get
	result, err := s.storage.GetClientState(s.key, s.clientType)
	s.NoError(err)
	s.Equal(clientState, result)

	// Check Update
	clientState = ratelimiter.ClientState{
		WindowStart:  now.Add(time.Second * 10),
		RequestCount: 20,
		Blocked:      true,
		BlockUntil:   now.Add(time.Second * 20),
	}

	err = s.storage.InsertOrUpdateClientState(s.key, clientState, s.clientType)
	s.NoError(err)

	// Check if Updated
	result, err = s.storage.GetClientState(s.key, s.clientType)
	s.NoError(err)
	s.Equal(clientState, result)

	// Check Delete
	err = s.storage.DeleteClientState(s.key, s.clientType)
	s.NoError(err)

	// Check if Deleted
	result, err = s.storage.GetClientState(s.key, s.clientType)
	s.Error(err)
	s.Equal(ratelimiter.ErrClientStateNotFound, err)
	s.Equal(ratelimiter.ClientState{}, result)
}
