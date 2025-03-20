package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/configs"
	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/internal/infra/cache"
	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/pkg/middleware/ratelimiter"
	"github.com/redis/go-redis/v9"

	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	apiKeyCache := cache.NewRedisApiKeyCache(redisClient)
	clientStateCache := cache.NewRedisClientStateCache(redisClient)

	// apiKeyCache := cache.NewInMemoryApiKeyCache()
	// clientStateCache := cache.NewInMemoryClientStateCache()

	for _, apiKey := range cfg.ApiKeysRateLimit {
		token := ratelimiter.ApiKey{
			Key:       apiKey.Key,
			RateLimit: apiKey.RateLimit,
		}

		err := apiKeyCache.InsertOrUpdateApiKey(token)
		if err != nil {
			log.Fatal(err)
		}
	}

	rateLimiterConfig := ratelimiter.RateLimiterConfig{
		IPRateLimit:      cfg.IpRateLimit,
		BlockTime:        time.Duration(cfg.BlockTimeSeconds) * time.Second,
		ClientStateCache: clientStateCache,
		ApiKeyCache:      apiKeyCache,
	}

	limiter := ratelimiter.NewRateLimiter(rateLimiterConfig)
	midleware := ratelimiter.NewRateLimiterMiddleware(limiter)

	r.Use(chi_middleware.RealIP)
	r.Use(midleware.Handler)
	r.Use(chi_middleware.Logger)
	r.Use(chi_middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Println("Server runnning at ", cfg.WebServerPort)
	http.ListenAndServe(":"+cfg.WebServerPort, r)
}
