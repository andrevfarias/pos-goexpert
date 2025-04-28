package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WeatherApiService struct {
	BaseURL string
	APIKey  string
}

type WeatherAPIResponse struct {
	Current struct {
		Temperature float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherApiServiceOutputDto struct {
	TempC float64 `json:"temp_c"`
}

func NewWeatherApiService(baseURL string, apiKey string) *WeatherApiService {
	return &WeatherApiService{BaseURL: baseURL, APIKey: apiKey}
}

func (s *WeatherApiService) GetWeatherByCity(ctx context.Context, city string) (*WeatherApiServiceOutputDto, error) {
	// Buscar a temperatura atual na localidade via WeatherAPI sem propagar o contexto de tracing
	escapedCity := url.QueryEscape(city)
	requestURI := fmt.Sprintf("%s/current.json?key=%s&q=%s", s.BaseURL, s.APIKey, escapedCity)

	// Não propagar o contexto, apenas fazer a chamada normal sem tracing
	weatherResponse, err := http.Get(requestURI)
	if err != nil {
		return nil, err
	}
	defer weatherResponse.Body.Close()

	// Verificar se a resposta é 200
	if weatherResponse.StatusCode != http.StatusOK {
		return nil, ErrApiError
	}

	// Ler o corpo da resposta
	weatherBody, err := io.ReadAll(weatherResponse.Body)
	if err != nil {
		return nil, err
	}

	// Parsear o corpo da resposta
	var weatherAPIResponse WeatherAPIResponse
	if err := json.Unmarshal(weatherBody, &weatherAPIResponse); err != nil {
		return nil, err
	}

	return &WeatherApiServiceOutputDto{
		TempC: weatherAPIResponse.Current.Temperature,
	}, nil
}
