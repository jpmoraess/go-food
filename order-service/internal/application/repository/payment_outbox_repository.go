package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/enum"
)

type PaymentOutboxRepository interface {
	Save(ctx context.Context, paymentOutbox *dto.PaymentOutbox) error
	FindByTypeAndSagaIdAndSagaStatus(ctx context.Context, outboxType string, sagaId uuid.UUID, SagaStatus ...enum.SagaStatus) *dto.PaymentOutbox
	DeleteByTypeAndOutboxStatusAndSagaStatus(ctx context.Context, outboxType string, outboxStatus enum.OutboxStatus, SagaStatus ...enum.SagaStatus) error
	FindByTypeAndOutboxStatusAndSagaStatus(ctx context.Context, outboxType string, outboxStatus enum.OutboxStatus, SagaStatus ...enum.SagaStatus) []*dto.PaymentOutbox
}
