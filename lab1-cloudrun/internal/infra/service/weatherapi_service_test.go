package service

import (
	"testing"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/config"
	"github.com/stretchr/testify/suite"
)

type WeatherApiServiceTestSuite struct {
	suite.Suite
	weatherApiService *WeatherApiService
}

func (s *WeatherApiServiceTestSuite) SetupSuite() {
	cfg, err := config.LoadConfig("./../../../")
	s.Require().NoError(err)

	s.weatherApiService = NewWeatherApiService(cfg.WeatherAPIBaseURL, cfg.WeatherAPIKey)
}

func TestWeatherApiServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WeatherApiServiceTestSuite))
}

func (s *WeatherApiServiceTestSuite) TestGetWeatherByCity_ShouldReturnWeather() {
	// Arrange
	city := "São Paulo"

	// Act
	weather, err := s.weatherApiService.GetWeatherByCity(city)

	// Assert
	s.Require().NoError(err)
	s.NotNil(weather)
	s.NotNil(weather.TempC)
}

func (s *WeatherApiServiceTestSuite) TestGetWeatherByCity_ShouldReturnAPIError() {
	// Arrange
	city := "!@#$%^&*()" // Nome de cidade com caracteres especiais que deve causar erro na API

	// Act
	weather, err := s.weatherApiService.GetWeatherByCity(city)

	// Assert
	s.Require().Error(err)
	s.Nil(weather)
}
