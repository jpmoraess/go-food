package helper

import (
	"context"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/mapper"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type CreateOrderHelper struct {
	orderRepo    domain.OrderRepository
	orderService domain.OrderDomainService
}

func NewCreateOrderHelper(orderRepo domain.OrderRepository, orderService domain.OrderDomainService) *CreateOrderHelper {
	return &CreateOrderHelper{orderRepo: orderRepo, orderService: orderService}
}

func (h *CreateOrderHelper) PersistOrder(ctx context.Context, input *dto.CreateOrderInputDTO) (*domain.OrderCreatedEvent, error) {
	orderMapper := mapper.OrderMapper{}
	order := orderMapper.MapOrderInputToDomain(input)

	orderCreatedEvent, err := h.orderService.ValidateAndInitiateOrder(order)
	if err != nil {
		return nil, err
	}

	order, err = h.orderRepo.Save(ctx, orderCreatedEvent.GetOrder())
	if err != nil {
		return nil, err
	}

	return orderCreatedEvent, nil
}
