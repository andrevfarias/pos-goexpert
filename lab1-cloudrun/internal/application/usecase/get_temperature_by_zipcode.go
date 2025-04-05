package usecase

import (
	"context"
	"fmt"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/service"
)

// GetTemperatureByZipCode é o caso de uso para obter a temperatura a partir de um CEP
type GetTemperatureByZipCode struct {
	zipCodeFinder  service.ZipCodeFinder
	weatherService service.WeatherService
}

// NewGetTemperatureByZipCode cria uma nova instância do caso de uso
func NewGetTemperatureByZipCode(
	zipCodeFinder service.ZipCodeFinder,
	weatherService service.WeatherService,
) *GetTemperatureByZipCode {
	return &GetTemperatureByZipCode{
		zipCodeFinder:  zipCodeFinder,
		weatherService: weatherService,
	}
}

// Execute executa o caso de uso para obter a temperatura a partir de um CEP
func (uc *GetTemperatureByZipCode) Execute(ctx context.Context, zipCode string) (*entity.Weather, error) {
	// Busca o endereço a partir do CEP
	address, err := uc.zipCodeFinder.FindAddressByZipCode(ctx, zipCode)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar endereço: %w", err)
	}

	// Busca a temperatura a partir da cidade
	weather, err := uc.weatherService.GetWeatherByCity(ctx, address.City)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar temperatura: %w", err)
	}

	return weather, nil
}
