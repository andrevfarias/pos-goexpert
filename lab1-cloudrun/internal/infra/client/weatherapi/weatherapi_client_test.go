package weatherapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetWeatherByCity(t *testing.T) {
	tests := []struct {
		name           string
		city           string
		mockResponse   interface{}
		mockStatusCode int
		expectedError  string
		expectedResult *entity.Weather
	}{
		{
			name: "deve retornar temperatura quando cidade é válida",
			city: "São Paulo",
			mockResponse: weatherAPIResponse{
				Current: struct {
					TempC float64 `json:"temp_c"`
				}{
					TempC: 25.5,
				},
			},
			mockStatusCode: http.StatusOK,
			expectedError:  "",
			expectedResult: entity.NewWeather(25.5),
		},
		{
			name:           "deve retornar erro quando serviço retorna status diferente de 200",
			city:           "São Paulo",
			mockResponse:   nil,
			mockStatusCode: http.StatusInternalServerError,
			expectedError:  "erro ao consultar WeatherAPI: status 500",
			expectedResult: nil,
		},
		{
			name:           "deve retornar erro quando resposta é inválida",
			city:           "São Paulo",
			mockResponse:   "resposta inválida que não é um JSON",
			mockStatusCode: http.StatusOK,
			expectedError:  "json: cannot unmarshal string into Go value of type weatherapi.weatherAPIResponse",
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verificar se os parâmetros da query estão corretos
				query := r.URL.Query()
				assert.Equal(t, "test-api-key", query.Get("key"))
				assert.Equal(t, tt.city, query.Get("q"))

				w.WriteHeader(tt.mockStatusCode)
				if tt.mockResponse != nil {
					json.NewEncoder(w).Encode(tt.mockResponse)
				}
			}))
			defer server.Close()

			client := NewClient(server.URL, "test-api-key", 5)

			// Act
			result, err := client.GetWeatherByCity(context.Background(), tt.city)

			// Assert
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.TempC, result.TempC)
			}
		})
	}
}

func TestClient_GetWeatherByCity_Timeout(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-api-key", 1)

	// Act
	result, err := client.GetWeatherByCity(context.Background(), "São Paulo")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
	assert.Nil(t, result)
}

func TestNewClient(t *testing.T) {
	// Arrange
	baseURL := "http://example.com"
	apiKey := "test-api-key"
	timeoutSeconds := 10

	// Act
	client := NewClient(baseURL, apiKey, timeoutSeconds)

	// Assert
	assert.NotNil(t, client)
	assert.Equal(t, baseURL, client.BaseURL)
	assert.Equal(t, apiKey, client.APIKey)
	assert.Equal(t, time.Duration(timeoutSeconds)*time.Second, client.HTTPClient.Timeout)
}
