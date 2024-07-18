package helper

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type SagaHelper struct {
	orderRepo domain.OrderRepository
}

func NewSagaHelper(orderRepo domain.OrderRepository) *SagaHelper {
	return &SagaHelper{orderRepo: orderRepo}
}

func (s *SagaHelper) FindOrder(orderID uuid.UUID) (*domain.Order, error) {
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("order with id: %s could not be found\n", orderID))
	}
	return order, nil
}

func (s *SagaHelper) SaveOrder(order *domain.Order) (*domain.Order, error) {
	saved, err := s.orderRepo.Save(order)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("order with id: %s could not be saved\n", order.ID))
	}
	return saved, nil
}

func (s *SagaHelper) OrderStatusToSagaStatus(orderStatus domain.OrderStatus) saga.SagaStatus {
	switch orderStatus {
	case domain.PAID:
		return saga.PROCESSING
	case domain.APPROVED:
		return saga.SUCCEEDED
	case domain.CANCELLING:
		return saga.COMPENSATING
	case domain.CANCELLED:
		return saga.COMPENSATED
	default:
		return saga.STARTED
	}
}
