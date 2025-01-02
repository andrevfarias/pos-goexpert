package database

import (
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/entity"
	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (o *OrderRepositoryMock) Save(order *entity.Order) error {
	args := o.Called(order)
	return args.Error(0)
}

func (o *OrderRepositoryMock) FindAll() ([]entity.Order, error) {
	args := o.Called()
	return args.Get(0).([]entity.Order), args.Error(1)
}
