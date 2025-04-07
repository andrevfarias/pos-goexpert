package usecase

import (
	"context"
	"testing"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/config"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/entity"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/service"
	"github.com/stretchr/testify/suite"
)

type GetTemperatureByZipCodeTestSuite struct {
	suite.Suite
	getTemperatureByZipCode *GetTemperatureByZipCode
}

func (s *GetTemperatureByZipCodeTestSuite) SetupTest() {
	cfg, err := config.LoadConfig("./../../")
	s.Require().NoError(err)

	viaCepService := service.NewViaCepApiService(cfg.ViacepAPIBaseURL)
	weatherService := service.NewWeatherApiService(cfg.WeatherAPIBaseURL, cfg.WeatherAPIKey)

	s.getTemperatureByZipCode = NewGetTemperatureByZipCode(viaCepService, weatherService)
}

func TestGetTemperatureByZipCodeTestSuite(t *testing.T) {
	suite.Run(t, new(GetTemperatureByZipCodeTestSuite))
}

func (s *GetTemperatureByZipCodeTestSuite) TestGetTemperatureByZipCode_ShouldReturnTemperature() {
	// Arrange (Valid Zipcode from Chapecó)
	zipCode, err := entity.NewZipCode("89802000")
	s.Require().NoError(err)

	// Act
	temperature, err := s.getTemperatureByZipCode.Execute(context.Background(), zipCode)

	// Assert
	s.Require().NoError(err)
	s.Require().NotNil(temperature)
}

func (s *GetTemperatureByZipCodeTestSuite) TestGetTemperatureByZipCode_ShouldReturnError() {
	// Arrange (Inexistent Zipcode)
	zipCode, err := entity.NewZipCode("12345678")
	s.Require().NoError(err)

	// Act
	temperature, err := s.getTemperatureByZipCode.Execute(context.Background(), zipCode)

	// Assert
	s.Require().Error(err)
	s.Equal(err, ErrZipCodeNotFound)
	s.Require().Nil(temperature)
}
