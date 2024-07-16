package messaging

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/gateway"
	"log"
)

type PaymentResponseKafka struct {
	listener gateway.PaymentMessageListener
}

func NewPaymentResponseKafka(listener gateway.PaymentMessageListener) *PaymentResponseKafka {
	return &PaymentResponseKafka{listener: listener}
}

func (p *PaymentResponseKafka) Receive() {
	// TODO: implement kafka consumer

	var err error
	err = p.listener.PaymentCompleted(&dto.PaymentResponse{})
	if err != nil {
		log.Println("error...")
	}

	err = p.listener.PaymentCancelled(&dto.PaymentResponse{})
	if err != nil {
		log.Println("error...")
	}
}
