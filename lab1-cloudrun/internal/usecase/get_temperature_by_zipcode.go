package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/entity"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/service"
)

var (
	ErrInvalidZipCode  = errors.New("invalid zipcode")
	ErrZipCodeNotFound = errors.New("can not find zipcode")
)

type GetTemperatureByZipCodeOutputDto struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

// GetTemperatureByZipCode é o caso de uso para obter a temperatura a partir de um CEP
type GetTemperatureByZipCode struct {
	viaCepService  *service.ViaCepApiService
	weatherService *service.WeatherApiService
}

// NewGetTemperatureByZipCode cria uma nova instância do caso de uso
func NewGetTemperatureByZipCode(
	viaCepService *service.ViaCepApiService,
	weatherService *service.WeatherApiService,
) *GetTemperatureByZipCode {
	return &GetTemperatureByZipCode{
		viaCepService:  viaCepService,
		weatherService: weatherService,
	}
}

// Execute executa o caso de uso para obter a temperatura a partir de um CEP
func (uc *GetTemperatureByZipCode) Execute(ctx context.Context, zipCode *entity.ZipCode) (*GetTemperatureByZipCodeOutputDto, error) {
	// Busca o endereço a partir do CEP
	address, err := uc.viaCepService.GetAddressByZipcode(zipCode.ZipCode)
	if err != nil {
		switch err {
		case service.ErrZipCodeNotFound:
			return nil, ErrZipCodeNotFound
		default:
			return nil, err
		}
	}

	// Busca a temperatura a partir da cidade
	weather, err := uc.weatherService.GetWeatherByCity(address.City)
	if err != nil {
		return nil, err
	}

	tempC := math.Round(weather.TempC*100) / 100
	tempF := math.Round((tempC*1.8+32)*100) / 100
	tempK := math.Round((tempC+273.15)*100) / 100

	return &GetTemperatureByZipCodeOutputDto{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}
