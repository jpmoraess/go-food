package helper

import (
	"context"

	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/mapper"
	"github.com/jpmoraess/go-food/order-service/internal/application/repository"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type CreateOrderHelper struct {
	orderMapper        *mapper.OrderMapper
	orderRepo          repository.OrderRepository
	orderDomainService domain.OrderDomainService
}

func NewCreateOrderHelper(orderMapper *mapper.OrderMapper, orderRepo repository.OrderRepository, orderDomainService domain.OrderDomainService) *CreateOrderHelper {
	return &CreateOrderHelper{orderMapper: orderMapper, orderRepo: orderRepo, orderDomainService: orderDomainService}
}

func (h *CreateOrderHelper) PersistOrder(ctx context.Context, input *dto.CreateOrderInputDTO) (*domain.OrderCreatedEvent, error) {
	order := h.orderMapper.CreateOrderInputToOrder(input)
	orderCreatedEvent, err := h.orderDomainService.InitiateOrder(order)
	if err != nil {
		return nil, err
	}

	_, err = h.orderRepo.Save(ctx, orderCreatedEvent.Order())
	if err != nil {
		return nil, err
	}

	return orderCreatedEvent, nil
}
