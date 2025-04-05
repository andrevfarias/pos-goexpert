package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
)

// Client é responsável por interagir com a API ViaCEP
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// viaCEPResponse representa a estrutura de resposta da API ViaCEP
type viaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Erro        bool   `json:"erro"`
}

// NewClient cria uma nova instância do cliente ViaCEP
func NewClient(baseURL string, timeoutSeconds int) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Duration(timeoutSeconds) * time.Second,
		},
	}
}

// GetAddressByZipCode consulta a API ViaCEP e retorna o endereço correspondente ao CEP fornecido
func (c *Client) GetAddressByZipCode(ctx context.Context, zipCode string) (*entity.Address, error) {
	url := fmt.Sprintf("%s/%s/json/", c.BaseURL, zipCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao consultar ViaCEP: status %d", resp.StatusCode)
	}

	var viaCEPResp viaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
		return nil, err
	}

	// Se a API retornar erro, retornamos erro
	if viaCEPResp.Erro {
		return nil, fmt.Errorf("CEP não encontrado")
	}

	// Converter a resposta da API para a entidade de domínio usando o construtor
	return entity.NewAddress(
		viaCEPResp.CEP,
		viaCEPResp.Logradouro,
		viaCEPResp.Complemento,
		viaCEPResp.Bairro,
		viaCEPResp.Localidade,
		viaCEPResp.UF,
	)
}
