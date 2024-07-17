package gateway

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"log"
)

type PaymentMessageListener interface {
	PaymentCompleted(*dto.PaymentResponse) error
	PaymentCancelled(*dto.PaymentResponse) error
}

type PaymentMessageListenerImpl struct {
	orderPaymentSaga *saga.OrderPaymentSaga
}

func NewPaymentMessageListenerImpl(orderPaymentSaga *saga.OrderPaymentSaga) *PaymentMessageListenerImpl {
	return &PaymentMessageListenerImpl{orderPaymentSaga: orderPaymentSaga}
}

func (p *PaymentMessageListenerImpl) PaymentCompleted(response *dto.PaymentResponse) error {
	err := p.orderPaymentSaga.Process(response)
	if err != nil {
		return err
	}
	log.Printf("order payment saga process is completed for order id: %s\n", response.OrderID)
	return nil
}

func (p *PaymentMessageListenerImpl) PaymentCancelled(response *dto.PaymentResponse) error {
	err := p.orderPaymentSaga.Rollback(response)
	if err != nil {
		return err
	}
	log.Printf("order is roll backed for order id: %s with failure messages: %s\n", response.OrderID, response.FailureMessages)
	return nil
}
