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
	saga saga.OrderPaymentSaga
}

func NewPaymentMessageListenerImpl(saga saga.OrderPaymentSaga) *PaymentMessageListenerImpl {
	return &PaymentMessageListenerImpl{saga: saga}
}

func (p *PaymentMessageListenerImpl) PaymentCompleted(response *dto.PaymentResponse) error {
	err := p.saga.Process(response)
	if err != nil {
		return err
	}
	log.Printf("order payment saga process is completed for order id: %s\n", response.OrderID)
	return nil
}

func (p *PaymentMessageListenerImpl) PaymentCancelled(response *dto.PaymentResponse) error {
	return nil
}
