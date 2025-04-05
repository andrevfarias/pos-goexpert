package main

import (
	"log"
	"net/http"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/application/usecase"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/config"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/api"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/api/handler"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/client/viacep"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/client/weatherapi"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/service"
)

func main() {
	// Carregar as configurações
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Inicializar os clientes HTTP
	viaCEPClient := viacep.NewClient(cfg.ViacepAPIBaseURL, cfg.APITimeoutSeconds)
	weatherAPIClient := weatherapi.NewClient(cfg.WeatherAPIBaseURL, cfg.WeatherAPIKey, cfg.APITimeoutSeconds)

	// Inicializar os serviços
	zipCodeFinderService := service.NewZipCodeFinderService(viaCEPClient)
	weatherService := service.NewWeatherService(weatherAPIClient)

	// Inicializar o caso de uso
	getTemperatureUseCase := usecase.NewGetTemperatureByZipCode(zipCodeFinderService, weatherService)

	// Inicializar o handler
	temperatureHandler := handler.NewTemperatureHandler(getTemperatureUseCase)

	// Inicializar o router
	router := api.NewRouter(temperatureHandler)

	// Iniciar o servidor
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando na porta %s...\n", port)
	log.Printf("Configurações: URL ViaCEP=%s, URL WeatherAPI=%s, Timeout=%d\n",
		cfg.ViacepAPIBaseURL, cfg.WeatherAPIBaseURL, cfg.APITimeoutSeconds)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
