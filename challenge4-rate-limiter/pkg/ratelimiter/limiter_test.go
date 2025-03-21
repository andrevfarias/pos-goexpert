package ratelimiter_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/storage/memory"
	"github.com/stretchr/testify/suite"
)

type LimiterTestSuite struct {
	suite.Suite
	ipRateLimit        int
	apiKeys            []ratelimiter.ApiKey
	blockTime          time.Duration
	clientStateStorage *memory.MemoryClientStateStorage
	apiKeyStorage      *memory.MemoryApiKeyStorage
	limiter            ratelimiter.RateLimiterService
}

func (s *LimiterTestSuite) SetupTest() {
	s.ipRateLimit = 5
	s.blockTime = time.Second * 10
	s.clientStateStorage = memory.NewMemoryClientStateStorage()
	s.apiKeyStorage = memory.NewMemoryApiKeyStorage()

	for i := range 2 {
		s.apiKeys = append(s.apiKeys, ratelimiter.ApiKey{
			Key:       fmt.Sprintf("key%d", i),
			RateLimit: 6,
		})
		s.apiKeyStorage.InsertOrUpdateApiKey(s.apiKeys[i])
	}

	s.limiter = ratelimiter.NewRateLimiter(ratelimiter.RateLimiterConfig{
		IPRateLimit:        s.ipRateLimit,
		BlockTime:          s.blockTime,
		ClientStateStorage: s.clientStateStorage,
		ApiKeyStorage:      s.apiKeyStorage,
	})
}

func TestLimiter(t *testing.T) {
	suite.Run(t, new(LimiterTestSuite))
}

func (s *LimiterTestSuite) TestIsIpAllowed() {
	// Check Empty IP
	allowed, err := s.limiter.IsIpAllowed("")
	s.Error(err)
	s.False(allowed)

	// Check Invalid IP
	allowed, err = s.limiter.IsIpAllowed("127.0.01")
	s.Error(err)
	s.False(allowed)

	// Check Valid IP
	allowed, err = s.limiter.IsIpAllowed("127.0.0.1")
	s.NoError(err)
	s.True(allowed)
}

func (s *LimiterTestSuite) TestIsIPAllowedConcurrency() {
	now := time.Now()
	requests := s.ipRateLimit + 3
	sleepTime := time.Millisecond * 50
	wg := sync.WaitGroup{}
	wg.Add(requests)

	for i := range requests {
		go func() {
			defer wg.Done()
			allowed, err := s.limiter.IsIpAllowed("127.0.0.2")
			if i < s.ipRateLimit {
				s.NoError(err)
				s.True(allowed)
			} else {
				s.NoError(err)
				s.False(allowed)
			}
		}()
		time.Sleep(sleepTime)
	}
	wg.Wait()
	clientState, _ := s.clientStateStorage.GetClientState("127.0.0.2", ratelimiter.ClientTypes.IP)
	s.Equal(true, clientState.Blocked)
	s.Equal(s.ipRateLimit+1, clientState.RequestCount)
	s.WithinDuration(now.Add(s.blockTime), clientState.BlockUntil, time.Millisecond*sleepTime*time.Duration(s.ipRateLimit+1))
}

func (s *LimiterTestSuite) TestIsApiKeyAllowed() {
	// Check Empty API Key
	allowed, err := s.limiter.IsApiKeyAllowed("")
	s.NoError(err)
	s.False(allowed)

	// Check Invalid API Key
	allowed, err = s.limiter.IsApiKeyAllowed("invalid_key")
	s.NoError(err)
	s.False(allowed)

	// Check Valid API Key
	allowed, err = s.limiter.IsApiKeyAllowed(s.apiKeys[0].Key)
	s.NoError(err)
	s.True(allowed)
}

func (s *LimiterTestSuite) TestIsApiKeyAllowedConcurrency() {
	now := time.Now()
	apiKey := s.apiKeys[0]
	requests := apiKey.RateLimit + 3
	sleepTime := time.Millisecond * 50
	wg := sync.WaitGroup{}
	wg.Add(requests)

	for i := range requests {
		go func() {
			defer wg.Done()
			allowed, err := s.limiter.IsApiKeyAllowed(apiKey.Key)
			if i < apiKey.RateLimit {
				s.NoError(err)
				s.True(allowed)
			} else {
				s.NoError(err)
				s.False(allowed)
			}
		}()
		time.Sleep(sleepTime)
	}
	wg.Wait()
	clientState, _ := s.clientStateStorage.GetClientState(apiKey.Key, ratelimiter.ClientTypes.ApiKey)
	s.Equal(true, clientState.Blocked)
	s.Equal(apiKey.RateLimit+1, clientState.RequestCount)
	s.WithinDuration(now.Add(s.blockTime), clientState.BlockUntil, time.Millisecond*sleepTime*time.Duration(apiKey.RateLimit+1))
	allowed, err := s.limiter.IsApiKeyAllowed(s.apiKeys[1].Key)
	s.NoError(err)
	s.True(allowed)
}
