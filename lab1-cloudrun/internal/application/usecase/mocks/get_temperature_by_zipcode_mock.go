package mocks

import (
	"context"

	"github.com/andrevfarias/go-expert/lab1-cloudrun/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

// GetTemperatureByZipCodeMock é um mock do use case GetTemperatureByZipCode
type GetTemperatureByZipCodeMock struct {
	mock.Mock
}

// Execute implementa a interface do use case
func (m *GetTemperatureByZipCodeMock) Execute(ctx context.Context, zipCode string) (*entity.Weather, error) {
	args := m.Called(ctx, zipCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Weather), args.Error(1)
}
