package saga

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/enum"
	"github.com/jpmoraess/go-food/order-service/internal/application/helper"
	"github.com/jpmoraess/go-food/order-service/internal/application/repository"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type OrderPaymentSaga struct {
	orderRepo           repository.OrderRepository
	domainService       domain.OrderDomainService
	paymentOutboxHelper *helper.PaymentOutboxHelper
}

func NewOrderPaymentSaga(
	orderRepo repository.OrderRepository,
	domainService domain.OrderDomainService,
	paymentOutboxHelper *helper.PaymentOutboxHelper,
) *OrderPaymentSaga {
	return &OrderPaymentSaga{
		orderRepo:           orderRepo,
		domainService:       domainService,
		paymentOutboxHelper: paymentOutboxHelper,
	}
}

func (s *OrderPaymentSaga) Process(response *dto.PaymentResponse) error {
	sagaID, _ := uuid.Parse(response.SagaID)

	paymentOutbox := s.paymentOutboxHelper.GetPaymentOutboxMessageBySagaIdAndSagaStatus(context.Background(), sagaID, enum.SAGA_STARTED)
	if paymentOutbox != nil {
		s.completePayment(response)
	}
	return nil
}

func (s *OrderPaymentSaga) Rollback(response *dto.PaymentResponse) error {
	log.Println("rollback orderPaymentSaga")
	return nil
}

func (s *OrderPaymentSaga) completePayment(response *dto.PaymentResponse) (*domain.OrderPaidEvent, error) {
	orderID, _ := uuid.Parse(response.OrderID)
	order, err := s.orderRepo.FindByID(context.Background(), orderID)
	if err != nil {
		return nil, err
	}
	orderPaidEvent, err := s.domainService.PayOrder(order)
	if err != nil {
		return nil, err
	}
	_, err = s.orderRepo.Save(context.Background(), orderPaidEvent.Order())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
