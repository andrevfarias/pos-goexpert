package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/andrevfarias/go-expert/lab2-observability/service-a/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type WebServerConfig struct {
	APITimeout         time.Duration
	RequestNameOTEL    string
	OTELCollectorURL   string
	OTELTracer         trace.Tracer
	ExternalAPIBaseURL string
}

type WebServer struct {
	apiTimeout         time.Duration
	requestNameOTEL    string
	otelCollectorURL   string
	otelTracer         trace.Tracer
	externalAPIBaseURL string
}

// NewServer creates a new WebServer
func NewServer(config *WebServerConfig) *WebServer {
	if err := validateConfig(config); err != nil {
		log.Fatal(err)
	}

	return &WebServer{
		apiTimeout:         config.APITimeout,
		requestNameOTEL:    config.RequestNameOTEL,
		otelCollectorURL:   config.OTELCollectorURL,
		otelTracer:         config.OTELTracer,
		externalAPIBaseURL: config.ExternalAPIBaseURL,
	}
}

func validateConfig(config *WebServerConfig) error {
	if config.OTELCollectorURL == "" {
		return fmt.Errorf("OTELCollectorURL is required")
	}
	if config.RequestNameOTEL == "" {
		return fmt.Errorf("RequestNameOTEL is required")
	}
	if config.OTELTracer == nil {
		return fmt.Errorf("OTELTracer is required")
	}
	if config.ExternalAPIBaseURL == "" {
		return fmt.Errorf("ExternalAPIBaseURL is required")
	}
	return nil
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
	router.Post("/cep", we.HandleRequest)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return router
}

func (ws *WebServer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Inicializa o contexto
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	// Inicia o span da requisição inicial
	ctx, span := ws.otelTracer.Start(ctx, ws.requestNameOTEL)
	defer span.End()

	// Parse JSON input instead of form data
	var requestData struct {
		CEP string `json:"cep"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	zipCode, err := entity.NewZipCode(requestData.CEP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	requestURI := fmt.Sprintf("%s/cep/%s", ws.externalAPIBaseURL, zipCode.ZipCode)

	req, err := http.NewRequestWithContext(ctx, "GET", requestURI, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Injeta o contexto no request
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Repassa o status HTTP correto da resposta do serviço B
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyBytes)
}
