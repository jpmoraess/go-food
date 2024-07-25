package helper

import (
	"context"

	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/mapper"
	"github.com/jpmoraess/go-food/order-service/internal/application/repository"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type CreateOrderHelper struct {
	orderRepo          repository.OrderRepository
	orderDomainService domain.OrderDomainService
}

func NewCreateOrderHelper(orderRepo repository.OrderRepository, orderDomainService domain.OrderDomainService) *CreateOrderHelper {
	return &CreateOrderHelper{orderRepo: orderRepo, orderDomainService: orderDomainService}
}

func (h *CreateOrderHelper) PersistOrder(ctx context.Context, input *dto.CreateOrderInputDTO) (*domain.OrderCreatedEvent, error) {
	orderMapper := mapper.OrderMapper{}
	order := orderMapper.CreateOrderInputToOrder(input)

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
