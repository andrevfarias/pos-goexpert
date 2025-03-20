package ratelimiter

import "errors"

var (
	ErrAPIKeyNotFound      = errors.New("API_KEY not found")
	ErrClientStateNotFound = errors.New("client state not found")
)
