package ratelimiter

import "time"

type ClientType string

const (
	clientTypeIP     ClientType = "ip"
	clientTypeApiKey ClientType = "apikey"
)

var ClientTypes = struct {
	IP     ClientType
	ApiKey ClientType
}{
	IP:     clientTypeIP,
	ApiKey: clientTypeApiKey,
}

type ApiKey struct {
	Key       string
	RateLimit int
}

type RateLimiterConfig struct {
	IPRateLimit        int
	BlockTime          time.Duration
	ClientStateStorage ClientStateStorage
	ApiKeyStorage      ApiKeyStorage
}

type ClientState struct {
	WindowStart  time.Time
	RequestCount int
	Blocked      bool
	BlockUntil   time.Time
}
