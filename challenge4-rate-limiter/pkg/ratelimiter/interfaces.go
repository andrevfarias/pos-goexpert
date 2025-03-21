package ratelimiter

type RateLimiterService interface {
	IsIpAllowed(ip string) (bool, error)
	IsApiKeyAllowed(apiKey string) (bool, error)
}
