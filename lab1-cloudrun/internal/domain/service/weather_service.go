package service

import (
	"context"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
)

// WeatherService define a interface para o serviço de busca de temperatura
type WeatherService interface {
	// GetWeatherByCity busca a temperatura atual para uma cidade
	GetWeatherByCity(ctx context.Context, city string) (*entity.Weather, error)
}
