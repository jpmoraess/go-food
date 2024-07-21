package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/go-food/order-service/internal/application/gateway"
	"github.com/jpmoraess/go-food/order-service/internal/application/helper"
	"github.com/jpmoraess/go-food/order-service/internal/application/outbox"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/application/usecase"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/db"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/handlers"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/message"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/persistence"
	"log"
)

func main() {
	dbpool, err := db.CreateConnectionPool()
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()

	orderRepository := persistence.NewOrderRepositoryPostgres(dbpool)
	paymentOutboxRepository := persistence.NewPaymentOutboxRepositoryPostgres(dbpool)

	paymentOutboxHelper := outbox.NewPaymentOutboxHelper(paymentOutboxRepository)

	orderDomainService := &domain.OrderDomainServiceImpl{}

	sagaHelper := helper.NewSagaHelper(orderRepository)
	createOrderHelper := helper.NewCreateOrderHelper(orderRepository, orderDomainService)

	orderPaymentSaga := saga.NewOrderPaymentSaga()
	paymentMessageListener := gateway.NewPaymentMessageListenerImpl(orderPaymentSaga)
	paymentResponseKafka := message.NewPaymentResponseKafka(paymentMessageListener)

	createOrderUseCase := usecase.NewCreateOrderUseCase(sagaHelper, createOrderHelper, paymentOutboxHelper)

	go paymentResponseKafka.StartConsume()

	// handlers
	orderHandler := handlers.NewOrderHandler(createOrderUseCase)

	// routes
	app := fiber.New()
	v1 := app.Group("/api/v1")
	v1.Post("/order", orderHandler.CreateOrder)

	log.Println("servidor HTTP rodando na porta 8080")
	if err = app.Listen(":8080"); err != nil {
		log.Fatalf("erro ao iniciar o servidor HTTP: %v\n", err)
	}
}
