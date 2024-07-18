package outbox

import (
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
	OrderID            string    `json:"orderId"`
	CustomerID         string    `json:"customerId"`
	Price              float64   `json:"price"`
	CreatedAt          time.Time `json:"createdAt"`
	PaymentOrderStatus string    `json:"paymentOrderStatus"`
}

type PaymentOutboxRepository interface {
	Save(paymentOutbox *PaymentOutbox) error
	FindByTypeAndSagaIdAndSagaStatus(outboxType string, sagaId uuid.UUID, SagaStatus ...saga.SagaStatus) *PaymentOutbox
	DeleteByTypeAndOutboxStatusAndSagaStatus(outboxType string, outboxStatus OutboxStatus, SagaStatus ...saga.SagaStatus)
	FindByTypeAndOutboxStatusAndSagaStatus(outboxType string, outboxStatus OutboxStatus, SagaStatus ...saga.SagaStatus) []*PaymentOutbox
}
