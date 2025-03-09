package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/configs"
	"github.com/andrevfarias/go-expert/challenge4-rate-limiter/middleware"

	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	rateLimiterConfig := middleware.RateLimiterConfig{
		IPRateLimit: cfg.IpRateLimit,
		APIKeyRateLimit: map[string]int{
			"key1": 10,
			"key2": 20,
			"key3": 30,
		},
		BlockTime: time.Duration(cfg.BlockTimeSeconds) * time.Second,
	}

	r.Use(chi_middleware.RealIP)
	r.Use(middleware.RateLimiter(rateLimiterConfig))
	r.Use(chi_middleware.Logger)
	r.Use(chi_middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Println("Server runnning at ", cfg.WebServerPort)
	http.ListenAndServe(":"+cfg.WebServerPort, r)
}
