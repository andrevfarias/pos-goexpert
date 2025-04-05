package viacep

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

func TestClient_GetAddressByZipCode(t *testing.T) {
	tests := []struct {
		name           string
		zipCode        string
		mockResponse   interface{}
		mockStatusCode int
		expectedError  string
		expectedResult *entity.Address
	}{
		{
			name:    "deve retornar endereço quando CEP é válido",
			zipCode: "12345678",
			mockResponse: viaCEPResponse{
				CEP:         "12345-678",
				Logradouro:  "Rua Teste",
				Complemento: "Casa 1",
				Bairro:      "Centro",
				Localidade:  "São Paulo",
				UF:          "SP",
				Erro:        false,
			},
			mockStatusCode: http.StatusOK,
			expectedError:  "",
			expectedResult: &entity.Address{
				ZipCode:      "12345-678",
				Street:       "Rua Teste",
				Complement:   "Casa 1",
				Neighborhood: "Centro",
				City:         "São Paulo",
				State:        "SP",
			},
		},
		{
			name:    "deve retornar erro quando CEP não é encontrado",
			zipCode: "99999999",
			mockResponse: viaCEPResponse{
				Erro: true,
			},
			mockStatusCode: http.StatusOK,
			expectedError:  "CEP não encontrado",
			expectedResult: nil,
		},
		{
			name:           "deve retornar erro quando serviço retorna status diferente de 200",
			zipCode:        "12345678",
			mockResponse:   nil,
			mockStatusCode: http.StatusInternalServerError,
			expectedError:  "erro ao consultar ViaCEP: status 500",
			expectedResult: nil,
		},
		{
			name:           "deve retornar erro quando resposta é inválida",
			zipCode:        "12345678",
			mockResponse:   "resposta inválida que não é um JSON",
			mockStatusCode: http.StatusOK,
			expectedError:  "json: cannot unmarshal string into Go value of type viacep.viaCEPResponse",
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				if tt.mockResponse != nil {
					json.NewEncoder(w).Encode(tt.mockResponse)
				}
			}))
			defer server.Close()

			client := NewClient(server.URL, 5)

			// Act
			result, err := client.GetAddressByZipCode(context.Background(), tt.zipCode)

			// Assert
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.ZipCode, result.ZipCode)
				assert.Equal(t, tt.expectedResult.Street, result.Street)
				assert.Equal(t, tt.expectedResult.Complement, result.Complement)
				assert.Equal(t, tt.expectedResult.Neighborhood, result.Neighborhood)
				assert.Equal(t, tt.expectedResult.City, result.City)
				assert.Equal(t, tt.expectedResult.State, result.State)
			}
		})
	}
}

func TestClient_GetAddressByZipCode_Timeout(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, 1)

	// Act
	result, err := client.GetAddressByZipCode(context.Background(), "12345678")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
	assert.Nil(t, result)
}

func TestNewClient(t *testing.T) {
	// Arrange
	baseURL := "http://example.com"
	timeoutSeconds := 10

	// Act
	client := NewClient(baseURL, timeoutSeconds)

	// Assert
	assert.NotNil(t, client)
	assert.Equal(t, baseURL, client.BaseURL)
	assert.Equal(t, time.Duration(timeoutSeconds)*time.Second, client.HTTPClient.Timeout)
}
