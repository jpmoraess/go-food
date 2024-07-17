package main

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/gateway"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/message"
)

func main() {
	orderPaymentSaga := saga.NewOrderPaymentSaga()
	paymentMessageListener := gateway.NewPaymentMessageListenerImpl(orderPaymentSaga)
	paymentResponseKafka := message.NewPaymentResponseKafka(paymentMessageListener)

	go paymentResponseKafka.StartConsume()
}
