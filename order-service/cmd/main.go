package main

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/gateway"
	"github.com/jpmoraess/go-food/order-service/internal/application/outbox"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/db"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/message"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/persistence"
)

func main() {
	pool, err := db.CreateConnectionPool()
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	paymentOutboxRepository := persistence.NewPaymentOutboxRepositoryPostgres(pool)
	_ = paymentOutboxRepository

	paymentOutboxHelper := outbox.NewPaymentOutboxHelper(paymentOutboxRepository)
	_ = paymentOutboxHelper

	orderPaymentSaga := saga.NewOrderPaymentSaga()
	paymentMessageListener := gateway.NewPaymentMessageListenerImpl(orderPaymentSaga)
	paymentResponseKafka := message.NewPaymentResponseKafka(paymentMessageListener)

	go paymentResponseKafka.StartConsume()
}
