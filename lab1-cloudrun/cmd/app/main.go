package main

import (
	"log"
	"net/http"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/config"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/api"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/api/handler"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/service"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/usecase"
)

func main() {
	// Carregar as configurações
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	viaCepService := service.NewViaCepApiService(cfg.ViacepAPIBaseURL)
	weatherService := service.NewWeatherApiService(cfg.WeatherAPIBaseURL, cfg.WeatherAPIKey)

	getTemperatureUseCase := usecase.NewGetTemperatureByZipCode(viaCepService, weatherService)

	temperatureHandler := handler.NewTemperatureHandler(getTemperatureUseCase)

	r := api.NewRouter(temperatureHandler)

	// Iniciar o servidor
	log.Printf("Servidor iniciando na porta %s...\n", cfg.Port)
	log.Printf("Configurações:\n\tURL ViaCEP=%s\n\tURL WeatherAPI=%s\n", cfg.ViacepAPIBaseURL, cfg.WeatherAPIBaseURL)

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
