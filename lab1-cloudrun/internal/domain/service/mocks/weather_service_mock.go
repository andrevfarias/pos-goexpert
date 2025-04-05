package mocks

import (
	"context"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

// WeatherServiceMock é um mock do WeatherService
type WeatherServiceMock struct {
	mock.Mock
}

// GetWeatherByCity implementa a interface WeatherService
func (m *WeatherServiceMock) GetWeatherByCity(ctx context.Context, city string) (*entity.Weather, error) {
	args := m.Called(ctx, city)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Weather), args.Error(1)
}
