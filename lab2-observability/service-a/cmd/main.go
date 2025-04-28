package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"

	"github.com/andrevfarias/go-expert/lab2-observability/service-a/internal/config"
	"github.com/andrevfarias/go-expert/lab2-observability/service-a/internal/infra/web"
)

func main() {
	// Carrega as configurações das variáveis de ambiente
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	// Exibe as configurações
	cfgJSON, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Printf("Erro ao converter configurações para JSON: %v", err)
	} else {
		log.Printf("Configurações carregadas:\n%s", string(cfgJSON))
	}

	// Trata graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Inicia o provider de tracing
	shutdown, err := initProvider(cfg.ServiceName, cfg.OTELCollectorURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TraceProvider: %w", err)
		}
	}()

	tracer := otel.Tracer(cfg.RequestNameOTEL)

	webServerConfig := &web.WebServerConfig{
		APITimeout:         time.Duration(cfg.APITimeoutSeconds) * time.Second,
		RequestNameOTEL:    cfg.RequestNameOTEL,
		OTELCollectorURL:   cfg.OTELCollectorURL,
		OTELTracer:         tracer,
		ExternalAPIBaseURL: cfg.ExternalAPIBaseURL,
	}

	server := web.NewServer(webServerConfig)
	router := server.CreateServer()

	go func() {
		log.Println("Starting server on port", cfg.Port)
		if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
			log.Fatal("failed to start server: %w", err)
		}
	}()

	select {
	case <-sigCh:
		log.Println("Received interrupt signal, shutting down gracefully...")
	case <-ctx.Done():
		log.Println("Context done, shutting down due to other reason...")
	}

	// Create a timeout context for the graceful shutdown
	_, shutdownCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer shutdownCancel()
}

func initProvider(serviceName string, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(collectorURL),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithTimeout(time.Second*10),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsb := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsb),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}
