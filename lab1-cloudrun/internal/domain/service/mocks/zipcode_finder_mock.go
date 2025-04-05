package mocks

import (
	"context"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

// ZipCodeFinderMock é um mock do ZipCodeFinder
type ZipCodeFinderMock struct {
	mock.Mock
}

// FindAddressByZipCode implementa a interface ZipCodeFinder
func (m *ZipCodeFinderMock) FindAddressByZipCode(ctx context.Context, zipCode string) (*entity.Address, error) {
	args := m.Called(ctx, zipCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Address), args.Error(1)
}
