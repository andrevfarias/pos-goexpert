package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/application/usecase/mocks"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureHandler_GetTemperatureByZipCode(t *testing.T) {
	tests := []struct {
		name             string
		zipCode          string
		setupMock        func(*mocks.GetTemperatureByZipCodeMock)
		expectedStatus   int
		expectedResponse interface{}
		expectedErrorMsg string
	}{
		{
			name:    "deve retornar temperatura quando CEP é válido",
			zipCode: "12345678",
			setupMock: func(m *mocks.GetTemperatureByZipCodeMock) {
				weather := entity.NewWeather(25.5)
				m.On("Execute", context.Background(), "12345678").Return(weather, nil)
			},
			expectedStatus: http.StatusOK,
			expectedResponse: entity.WeatherJSON{
				TempC: 25.5,
				TempF: 77.9,
				TempK: 298.65,
			},
		},
		{
			name:             "deve retornar erro quando CEP está vazio",
			zipCode:          "",
			setupMock:        func(m *mocks.GetTemperatureByZipCodeMock) {},
			expectedStatus:   http.StatusBadRequest,
			expectedErrorMsg: "zipcode query parameter is required",
		},
		{
			name:             "deve retornar erro quando CEP é inválido",
			zipCode:          "1234567",
			setupMock:        func(m *mocks.GetTemperatureByZipCodeMock) {},
			expectedStatus:   http.StatusUnprocessableEntity,
			expectedErrorMsg: "invalid zipcode",
		},
		{
			name:    "deve retornar erro quando CEP não é encontrado",
			zipCode: "12345678",
			setupMock: func(m *mocks.GetTemperatureByZipCodeMock) {
				m.On("Execute", context.Background(), "12345678").Return(nil, errors.New("CEP não encontrado"))
			},
			expectedStatus:   http.StatusNotFound,
			expectedErrorMsg: "can not find zipcode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			useCase := new(mocks.GetTemperatureByZipCodeMock)
			tt.setupMock(useCase)
			handler := NewTemperatureHandler(useCase)

			req := httptest.NewRequest(http.MethodGet, "/temperature?zipcode="+tt.zipCode, nil)
			w := httptest.NewRecorder()

			// Act
			handler.GetTemperatureByZipCode(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedErrorMsg != "" {
				assert.Equal(t, tt.expectedErrorMsg+"\n", w.Body.String())
			} else {
				var response entity.WeatherJSON
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NoError(t, err)
				assert.InDelta(t, tt.expectedResponse.(entity.WeatherJSON).TempC, response.TempC, 0.01)
				assert.InDelta(t, tt.expectedResponse.(entity.WeatherJSON).TempF, response.TempF, 0.1)
				assert.InDelta(t, tt.expectedResponse.(entity.WeatherJSON).TempK, response.TempK, 0.01)
			}

			// Verify mock
			useCase.AssertExpectations(t)
		})
	}
}
