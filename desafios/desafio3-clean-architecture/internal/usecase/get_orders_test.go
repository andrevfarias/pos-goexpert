package usecase

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/andrefarias66/pos-goexpert/desafios/desafio3-clean-architecture/internal/entity"
	"github.com/andrefarias66/pos-goexpert/desafios/desafio3-clean-architecture/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWhenGetOrders_ThenShouldReturnAllOrders(t *testing.T) {
	mockRepo := &database.OrderRepositoryMock{}

	mockOrders := []entity.Order{}
	for i := 1; i <= 10; i++ {
		order, _ := entity.NewOrder("order"+fmt.Sprint(i), float64(rand.IntN(20))+rand.Float64(), float64(rand.IntN(20))+rand.Float64())
		order.CalculateFinalPrice()
		mockOrders = append(mockOrders, *order)
	}

	mockRepo.On("FindAll", mock.Anything).Return([]entity.Order{}, nil).Once()
	mockRepo.On("FindAll", mock.Anything).Return(mockOrders, nil)

	usecase := NewGetOrdersUseCase(mockRepo)

	outputEmpty, err := usecase.Execute(GetOrdersInputDTO{})
	assert.Nil(t, err)
	assert.NotNil(t, outputEmpty)
	assert.NotNil(t, outputEmpty.Orders)
	assert.Equal(t, 0, len(outputEmpty.Orders))

	output, err := usecase.Execute(GetOrdersInputDTO{})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, 10, len(output.Orders))

	for i := 0; i < 10; i++ {
		assert.Equal(t, mockOrders[i].ID, output.Orders[i].ID)
		assert.Equal(t, mockOrders[i].Price, output.Orders[i].Price)
		assert.Equal(t, mockOrders[i].Tax, output.Orders[i].Tax)
		assert.Equal(t, mockOrders[i].FinalPrice, output.Orders[i].FinalPrice)
	}
}
