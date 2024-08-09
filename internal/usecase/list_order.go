package usecase

import (
	"github.com/nicolastanski/go-clean-architecture/internal/entity"
	"github.com/nicolastanski/go-clean-architecture/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {

	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	dto := make([]OrderOutputDTO, 0)
	for _, order := range orders {
		dto = append(dto, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	c.OrderListed.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderListed)

	return dto, nil
}
