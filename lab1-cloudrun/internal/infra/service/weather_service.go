package service

import (
	"context"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/client/weatherapi"
)

// WeatherServiceImpl implementa a interface WeatherService usando a API WeatherAPI
type WeatherServiceImpl struct {
	weatherAPIClient *weatherapi.Client
}

// NewWeatherService cria uma nova instância do serviço de busca de temperatura
func NewWeatherService(weatherAPIClient *weatherapi.Client) *WeatherServiceImpl {
	return &WeatherServiceImpl{
		weatherAPIClient: weatherAPIClient,
	}
}

// GetWeatherByCity busca a temperatura atual para uma cidade
func (s *WeatherServiceImpl) GetWeatherByCity(ctx context.Context, city string) (*entity.Weather, error) {
	return s.weatherAPIClient.GetWeatherByCity(ctx, city)
}
