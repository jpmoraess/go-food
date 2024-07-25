package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/go-food/order-service/internal/application/gateway"
	"github.com/jpmoraess/go-food/order-service/internal/application/helper"
	"github.com/jpmoraess/go-food/order-service/internal/application/mapper"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/application/usecase"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/db"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/handlers"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/messaging"
	"github.com/jpmoraess/go-food/order-service/internal/infrastructure/persistence"
)

func main() {
	// starting database connection pool
	dbpool, err := db.CreateConnectionPool()
	if err != nil {
		log.Fatalf("Erro ao criar pool de conexão com o banco de dados: %v\n", err)
	}
	defer dbpool.Close()

	// repositories
	orderRepo := persistence.NewOrderRepositoryPostgres(dbpool)
	paymentOutboxRepo := persistence.NewPaymentOutboxRepositoryPostgres(dbpool)

	// mappers
	orderMapper := &mapper.OrderMapper{}

	// services and helpers
	domainService := domain.NewOrderDomainServiceImpl()
	sagaHelper := &helper.SagaHelper{}
	createOrderHelper := helper.NewCreateOrderHelper(orderMapper, orderRepo, domainService)
	paymentOutboxHelper := helper.NewPaymentOutboxHelper(paymentOutboxRepo)

	// saga and messages
	orderPaymentSaga := saga.NewOrderPaymentSaga(orderRepo, domainService, paymentOutboxHelper)
	paymentMessageListener := gateway.NewPaymentMessageListenerImpl(orderPaymentSaga)
	paymentResponseKafka := messaging.NewPaymentResponseKafka(paymentMessageListener)

	// usecases
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderMapper, sagaHelper, createOrderHelper, paymentOutboxHelper)

	// http server
	app := fiber.New()
	setupRoutes(app, createOrderUseCase)

	// kafka consumer
	go paymentResponseKafka.StartConsume()

	// starting http server
	port := ":8090"
	log.Printf("HTTP server running on port %s\n", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting HTTP server: %v\n", err)
	}
}

func setupRoutes(app *fiber.App, createOrderUseCase *usecase.CreateOrderUseCase) {
	// handlers
	orderHandler := handlers.NewOrderHandler(createOrderUseCase)

	// routes
	v1 := app.Group("/api/v1")
	v1.Post("/order", orderHandler.CreateOrder)
}
