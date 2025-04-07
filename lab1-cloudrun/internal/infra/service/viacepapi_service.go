package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var ErrZipCodeNotFound = errors.New("zipcode not found")
var ErrApiError = errors.New("api error")

type ViaCepApiService struct {
	BaseURL string
}

type AddressDto struct {
	City string `json:"localidade"`
}

func NewViaCepApiService(baseURL string) *ViaCepApiService {
	return &ViaCepApiService{BaseURL: baseURL}
}

func (s *ViaCepApiService) GetAddressByZipcode(zipcode string) (*AddressDto, error) {
	// Pesquisar o CEP na API ViaCEP
	viaCepResponse, err := http.Get(fmt.Sprintf("%s/%s/json", s.BaseURL, zipcode))
	if err != nil {
		return nil, ErrApiError
	}
	defer viaCepResponse.Body.Close()

	// Verificar se a resposta é 200
	if viaCepResponse.StatusCode != http.StatusOK {
		return nil, ErrApiError
	}

	// Ler o corpo da resposta
	viaCepBody, err := io.ReadAll(viaCepResponse.Body)
	if err != nil {
		return nil, ErrApiError
	}

	// Verificar se retornou erro (cep não encontrado). Body: "{\n  \"erro\": \"true\"\n}"
	if strings.Contains(string(viaCepBody), "\"erro\": \"true\"") {
		return nil, ErrZipCodeNotFound
	}

	// Parsear o corpo da resposta
	var addressDto AddressDto
	if err := json.Unmarshal(viaCepBody, &addressDto); err != nil {
		return nil, ErrApiError
	}

	return &addressDto, nil
}
