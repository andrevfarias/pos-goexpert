//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/entity"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/event"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/infra/database"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/infra/web"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/usecase"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewGetOrdersUseCase(db *sql.DB) *usecase.GetOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewGetOrdersUseCase,
	)
	return &usecase.GetOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
