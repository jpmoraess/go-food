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

func (s *SagaHelper) findOrder(orderID uuid.UUID) (*domain.Order, error) {
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("order with id: %s could not be found\n", orderID))
	}
	return order, nil
}

func (s *SagaHelper) saveOrder(order *domain.Order) (*domain.Order, error) {
	saved, err := s.orderRepo.Save(order)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("order with id: %s could not be saved\n", order.ID))
	}
	return saved, nil
}

func orderStatusToSagaStatus(orderStatus domain.OrderStatus) saga.SagaStatus {
	switch orderStatus {
	case domain.Paid:
		return saga.PROCESSING
	case domain.Approved:
		return saga.SUCCEEDED
	case domain.Cancelling:
		return saga.COMPENSATING
	case domain.Cancelled:
		return saga.COMPENSATED
	default:
		return saga.STARTED
	}
}
