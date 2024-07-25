package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/enum"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type PaymentOutbox struct {
	ID           uuid.UUID          `json:"id"`
	SagaID       uuid.UUID          `json:"sagaId"`
	CreatedAt    time.Time          `json:"createdAt"`
	ProcessedAt  time.Time          `json:"processedAt"`
	Type         string             `json:"type"`
	Payload      string             `json:"payload"`
	OrderStatus  domain.OrderStatus `json:"orderStatus"`
	SagaStatus   enum.SagaStatus    `json:"sagaStatus"`
	OutboxStatus enum.OutboxStatus  `json:"outboxStatus"`
	Version      int                `json:"version"`
}

type PaymentEventPayload struct {
	OrderID            string                    `json:"orderId"`
	CustomerID         string                    `json:"customerId"`
	Price              float64                   `json:"price"`
	CreatedAt          time.Time                 `json:"createdAt"`
	PaymentOrderStatus domain.PaymentOrderStatus `json:"paymentOrderStatus"`
}
