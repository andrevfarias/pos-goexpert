package memory

import (
	"testing"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorWhenApiKeyIsNotFound(t *testing.T) {
	storage := NewMemoryApiKeyStorage()
	_, err := storage.GetApiKey("test")
	assert.Error(t, err)
}

func TestShouldReturnApiKeyWhenItExists(t *testing.T) {
	storage := NewMemoryApiKeyStorage()
	storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 10})
	apiKey, err := storage.GetApiKey("test")
	assert.NoError(t, err)
	assert.Equal(t, "test", apiKey.Key)
	assert.Equal(t, 10, apiKey.RateLimit)
}

func TestShouldUpdateApiKeyWhenItExists(t *testing.T) {
	storage := NewMemoryApiKeyStorage()
	storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 10})
	storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 20})
	apiKey, err := storage.GetApiKey("test")
	assert.NoError(t, err)
	assert.Equal(t, "test", apiKey.Key)
	assert.Equal(t, 20, apiKey.RateLimit)
}

func TestShouldDeleteApiKeyWhenItExists(t *testing.T) {
	storage := NewMemoryApiKeyStorage()
	storage.InsertOrUpdateApiKey(ratelimiter.ApiKey{Key: "test", RateLimit: 10})
	storage.DeleteApiKey("test")
	_, err := storage.GetApiKey("test")
	assert.Error(t, err)
}
