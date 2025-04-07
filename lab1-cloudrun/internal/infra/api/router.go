package api

import (
	"net/http"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter cria um novo roteador HTTP para a API
func NewRouter(temperatureHandler *handler.TemperatureHandler) http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rotas
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Get("/cep/{cep}", temperatureHandler.GetTemperatureByZipCode)

	return r
}
