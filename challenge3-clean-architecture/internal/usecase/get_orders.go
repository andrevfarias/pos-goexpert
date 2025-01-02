package usecase

import (
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/entity"
)

type GetOrdersInputDTO struct{}
type GetOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrdersUseCase) Execute(input GetOrdersInputDTO) (GetOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return GetOrdersOutputDTO{}, err
	}

	dto := GetOrdersOutputDTO{
		Orders: []OrderOutputDTO{},
	}

	for _, order := range orders {
		dto.Orders = append(dto.Orders, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		})
	}

	return dto, nil
}
