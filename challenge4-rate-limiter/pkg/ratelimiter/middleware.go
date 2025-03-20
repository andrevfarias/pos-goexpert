package ratelimiter

import (
	"net"
	"net/http"
)

type RateLimiterMiddleware struct {
	limiter RateLimiterService
}

func NewRateLimiterMiddleware(limiter RateLimiterService) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		limiter: limiter,
	}
}

func (rl *RateLimiterMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiKey := r.Header.Get("API_KEY"); apiKey != "" {
			isAllowed, err := rl.limiter.IsApiKeyAllowed(apiKey)

			if err == nil && !isAllowed {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			if err == nil && isAllowed {
				next.ServeHTTP(w, r)
				return
			}

			if err != nil && err != ErrAPIKeyNotFound {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if (err != nil) || (ip == "") {
			http.Error(w, "cannot resolve origin", http.StatusInternalServerError)
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
