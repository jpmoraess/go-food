package outbox

import (
	"context"
	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
	"time"
)

type PaymentOutbox struct {
	ID           uuid.UUID          `json:"id"`
	SagaID       uuid.UUID          `json:"sagaId"`
	CreatedAt    time.Time          `json:"createdAt"`
	ProcessedAt  time.Time          `json:"processedAt"`
	Type         string             `json:"type"`
	Payload      string             `json:"payload"`
	OrderStatus  domain.OrderStatus `json:"orderStatus"`
	SagaStatus   saga.SagaStatus    `json:"sagaStatus"`
	OutboxStatus OutboxStatus       `json:"outboxStatus"`
	Version      int                `json:"version"`
}

type PaymentEventPayload struct {
	OrderID            string                    `json:"orderId"`
	CustomerID         string                    `json:"customerId"`
	Price              float64                   `json:"price"`
	CreatedAt          time.Time                 `json:"createdAt"`
	PaymentOrderStatus domain.PaymentOrderStatus `json:"paymentOrderStatus"`
}

type PaymentOutboxRepository interface {
	Save(ctx context.Context, paymentOutbox *PaymentOutbox) error
	FindByTypeAndSagaIdAndSagaStatus(ctx context.Context, outboxType string, sagaId uuid.UUID, SagaStatus ...saga.SagaStatus) *PaymentOutbox
	DeleteByTypeAndOutboxStatusAndSagaStatus(ctx context.Context, outboxType string, outboxStatus OutboxStatus, SagaStatus ...saga.SagaStatus) error
	FindByTypeAndOutboxStatusAndSagaStatus(ctx context.Context, outboxType string, outboxStatus OutboxStatus, SagaStatus ...saga.SagaStatus) []*PaymentOutbox
}
