package service

import (
	"context"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/infra/client/viacep"
)

// ZipCodeFinderService implementa a interface ZipCodeFinder usando a API ViaCEP
type ZipCodeFinderService struct {
	viaCEPClient *viacep.Client
}

// NewZipCodeFinderService cria uma nova instância do serviço de busca de endereço por CEP
func NewZipCodeFinderService(viaCEPClient *viacep.Client) *ZipCodeFinderService {
	return &ZipCodeFinderService{
		viaCEPClient: viaCEPClient,
	}
}

// FindAddressByZipCode busca um endereço usando a API ViaCEP
func (s *ZipCodeFinderService) FindAddressByZipCode(ctx context.Context, zipCode string) (*entity.Address, error) {
	return s.viaCEPClient.GetAddressByZipCode(ctx, zipCode)
}
