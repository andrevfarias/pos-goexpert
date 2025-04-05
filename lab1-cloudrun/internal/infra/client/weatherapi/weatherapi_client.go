package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
)

// Client é responsável por interagir com a API WeatherAPI
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// weatherAPIResponse representa a estrutura de resposta da API WeatherAPI
type weatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

// NewClient cria uma nova instância do cliente WeatherAPI
func NewClient(baseURL, apiKey string, timeoutSeconds int) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Duration(timeoutSeconds) * time.Second,
		},
	}
}

// GetWeatherByCity consulta a API WeatherAPI e retorna a temperatura atual da cidade fornecida
func (c *Client) GetWeatherByCity(ctx context.Context, city string) (*entity.Weather, error) {
	// Construir a URL com query parameters
	requestURL, err := url.Parse(fmt.Sprintf("%s/current.json", c.BaseURL))
	if err != nil {
		return nil, err
	}

	// Adicionar query parameters
	query := requestURL.Query()
	query.Add("key", c.APIKey)
	query.Add("q", city)
	requestURL.RawQuery = query.Encode()

	// Criar a requisição
	req, err := http.NewRequestWithContext(ctx, "GET", requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Executar a requisição
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verificar o status da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao consultar WeatherAPI: status %d", resp.StatusCode)
	}

	// Decodificar a resposta
	var weatherResp weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}

	// Converter para Fahrenheit: F = C * 9/5 + 32
	tempF := weatherResp.Current.TempC*9.0/5.0 + 32.0

	// Converter para Kelvin: K = C + 273.15
	tempK := weatherResp.Current.TempC + 273.15

	// Criar a entidade de domínio
	weather := &entity.Weather{
		TempC: weatherResp.Current.TempC,
		TempF: tempF,
		TempK: tempK,
	}

	return weather, nil
}
