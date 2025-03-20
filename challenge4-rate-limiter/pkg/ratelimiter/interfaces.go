package ratelimiter

type ClientStateCache interface {
	GetClientState(clientID string, clientType ClientType) (ClientState, error)
	InsertOrUpdateClientState(clientID string, clientState ClientState, clientType ClientType) error
	DeleteClientState(clientID string, clientType ClientType) error
}

type ApiKeyCache interface {
	GetApiKey(apiKey string) (ApiKey, error)
	InsertOrUpdateApiKey(apiKey ApiKey) error
	DeleteApiKey(apiKey string) error
}

type RateLimiterService interface {
	IsIpAllowed(ip string) (bool, error)
	IsApiKeyAllowed(apiKey string) (bool, error)
}
