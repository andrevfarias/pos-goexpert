package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTemperatureByZipCode_Execute(t *testing.T) {
	tests := []struct {
		name           string
		zipCode        string
		setupMocks     func(*mocks.ZipCodeFinderMock, *mocks.WeatherServiceMock)
		expectedResult *entity.Weather
		expectedError  error
	}{
		{
			name:    "deve retornar temperatura quando CEP é válido",
			zipCode: "12345678",
			setupMocks: func(zf *mocks.ZipCodeFinderMock, ws *mocks.WeatherServiceMock) {
				address := &entity.Address{
					ZipCode: "12345678",
					City:    "São Paulo",
					State:   "SP",
				}
				weather := entity.NewWeather(25.5)

				zf.On("FindAddressByZipCode", mock.Anything, "12345678").Return(address, nil)
				ws.On("GetWeatherByCity", mock.Anything, "São Paulo").Return(weather, nil)
			},
			expectedResult: entity.NewWeather(25.5),
			expectedError:  nil,
		},
		{
			name:    "deve retornar erro quando CEP é inválido",
			zipCode: "123456",
			setupMocks: func(zf *mocks.ZipCodeFinderMock, ws *mocks.WeatherServiceMock) {
				zf.On("FindAddressByZipCode", mock.Anything, "123456").Return(nil, entity.ErrInvalidZipCode)
			},
			expectedResult: nil,
			expectedError:  errors.New("erro ao buscar endereço: CEP inválido"),
		},
		{
			name:    "deve retornar erro quando serviço de CEP está indisponível",
			zipCode: "12345678",
			setupMocks: func(zf *mocks.ZipCodeFinderMock, ws *mocks.WeatherServiceMock) {
				zf.On("FindAddressByZipCode", mock.Anything, "12345678").Return(nil, errors.New("serviço indisponível"))
			},
			expectedResult: nil,
			expectedError:  errors.New("erro ao buscar endereço: serviço indisponível"),
		},
		{
			name:    "deve retornar erro quando serviço de temperatura está indisponível",
			zipCode: "12345678",
			setupMocks: func(zf *mocks.ZipCodeFinderMock, ws *mocks.WeatherServiceMock) {
				address := &entity.Address{
					ZipCode: "12345678",
					City:    "São Paulo",
					State:   "SP",
				}
				zf.On("FindAddressByZipCode", mock.Anything, "12345678").Return(address, nil)
				ws.On("GetWeatherByCity", mock.Anything, "São Paulo").Return(nil, errors.New("serviço indisponível"))
			},
			expectedResult: nil,
			expectedError:  errors.New("erro ao buscar temperatura: serviço indisponível"),
		},
		{
			name:    "deve retornar erro quando CEP está vazio",
			zipCode: "",
			setupMocks: func(zf *mocks.ZipCodeFinderMock, ws *mocks.WeatherServiceMock) {
				zf.On("FindAddressByZipCode", mock.Anything, "").Return(nil, entity.ErrInvalidZipCode)
			},
			expectedResult: nil,
			expectedError:  errors.New("erro ao buscar endereço: CEP inválido"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			zipCodeFinder := new(mocks.ZipCodeFinderMock)
			weatherService := new(mocks.WeatherServiceMock)
			tt.setupMocks(zipCodeFinder, weatherService)

			useCase := NewGetTemperatureByZipCode(zipCodeFinder, weatherService)

			// Act
			result, err := useCase.Execute(context.Background(), tt.zipCode)

			// Assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.TempC, result.TempC)
			}

			// Verify mocks
			zipCodeFinder.AssertExpectations(t)
			weatherService.AssertExpectations(t)
		})
	}
}
