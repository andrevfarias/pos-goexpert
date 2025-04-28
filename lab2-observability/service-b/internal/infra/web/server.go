package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andrevfarias/go-expert/lab2-observability/service-b/internal/entity"
	"github.com/andrevfarias/go-expert/lab2-observability/service-b/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type WebServerConfig struct {
	APITimeout            time.Duration
	RequestNameOTEL       string
	OTELCollectorURL      string
	OTELTracer            trace.Tracer
	GetTemperatureUseCase *usecase.GetTemperatureByZipCode
}

type WebServer struct {
	apiTimeout            time.Duration
	requestNameOTEL       string
	otelCollectorURL      string
	otelTracer            trace.Tracer
	getTemperatureUseCase *usecase.GetTemperatureByZipCode
}

// NewServer creates a new WebServer
func NewServer(config *WebServerConfig) *WebServer {
	return &WebServer{
		apiTimeout:            config.APITimeout,
		requestNameOTEL:       config.RequestNameOTEL,
		otelCollectorURL:      config.OTELCollectorURL,
		otelTracer:            config.OTELTracer,
		getTemperatureUseCase: config.GetTemperatureUseCase,
	}
}

// CreateServer creates a new server instance with go chi router
func (we *WebServer) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(we.apiTimeout))
	// promhttp
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/cep/{cep}", we.HandleRequest)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return router
}

type TemplateData struct {
	Title              string
	BackgroundColour   string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}

func (h *WebServer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := h.otelTracer.Start(ctx, h.requestNameOTEL)
	defer span.End()

	// Obter o CEP do caminho (path)
	rawZipCode := r.PathValue("cep")
	zipCode, err := entity.NewZipCode(rawZipCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	weather, err := h.getTemperatureUseCase.Execute(ctx, zipCode)
	if err != nil {
		switch err {
		case usecase.ErrInvalidZipCode:
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		case usecase.ErrZipCodeNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Retornar o resultado como JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weather); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
