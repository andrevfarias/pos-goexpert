package ratelimiter

type ClientStateStorage interface {
	GetClientState(clientID string, clientType ClientType) (ClientState, error)
	InsertOrUpdateClientState(clientID string, clientState ClientState, clientType ClientType) error
	DeleteClientState(clientID string, clientType ClientType) error
}

type ApiKeyStorage interface {
	GetApiKey(apiKey string) (ApiKey, error)
	InsertOrUpdateApiKey(apiKey ApiKey) error
	DeleteApiKey(apiKey string) error
}
