package service

import (
	"context"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
)

// ZipCodeFinder define a interface para o serviço de busca de endereço por CEP
type ZipCodeFinder interface {
	// FindAddressByZipCode busca um endereço com base no CEP fornecido
	FindAddressByZipCode(ctx context.Context, zipCode string) (*entity.Address, error)
}
