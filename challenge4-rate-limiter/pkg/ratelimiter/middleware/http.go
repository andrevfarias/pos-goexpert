package middleware

import (
	"net"
	"net/http"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/ratelimiter"
)

type RateLimiterMiddleware struct {
	limiter ratelimiter.RateLimiterService
}

func NewRateLimiterMiddleware(limiter ratelimiter.RateLimiterService) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		limiter: limiter,
	}
}

func (rl *RateLimiterMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiKey := r.Header.Get("API_KEY"); apiKey != "" {
			isAllowed, err := rl.limiter.IsApiKeyAllowed(apiKey)
			if err != nil && err != ratelimiter.ErrAPIKeyNotFound {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err == nil && !isAllowed {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}
			if err == nil && isAllowed {
				next.ServeHTTP(w, r)
				return
			}
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if (err != nil) || (ip == "") {
			http.Error(w, ratelimiter.ErrInvalidIP.Error(), http.StatusInternalServerError)
			return
		}

		isAllowed, err := rl.limiter.IsIpAllowed(ip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !isAllowed {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
