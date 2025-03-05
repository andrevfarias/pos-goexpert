package main

import (
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(RateLimiterMidleware)
	// r.Use(ratelimiter.NewRateLimiterMidleware())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Println("Server runnning at :8080")
	http.ListenAndServe(":8080", r)
}

func RateLimiterMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the API_KEY from the request header
		api_key := r.Header.Get("API_KEY")
		if api_key != "" {
			log.Println("API_KEY:", api_key)
			// Check if the API_KEY is allowed to access the endpoint
			// If the API_KEY is allowed, call the next handler
			// If the API_KEY is not allowed, return a 429 status code

			next.ServeHTTP(w, r)
			return
		}
		// check the IP address
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if (err != nil) || (ip == "") {
			http.Error(w, "cannot resolve origin", http.StatusInternalServerError)
			return
		}

		log.Println("IP Address:", ip)
		// Check if the IP address is allowed to access the endpoint

		// If the IP address is allowed, call the next handler
		// If the IP address is not allowed, return a 429 status code
		next.ServeHTTP(w, r)
	})
}
