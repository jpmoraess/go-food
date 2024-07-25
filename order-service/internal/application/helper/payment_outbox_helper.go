package helper

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/enum"
	"github.com/jpmoraess/go-food/order-service/internal/application/repository"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type PaymentOutboxHelper struct {
	paymentOutboxRepo repository.PaymentOutboxRepository
}

func NewPaymentOutboxHelper(paymentOutboxRepo repository.PaymentOutboxRepository) *PaymentOutboxHelper {
	return &PaymentOutboxHelper{paymentOutboxRepo: paymentOutboxRepo}
}

func (h *PaymentOutboxHelper) GetPaymentOutboxByOutboxStatusAndSagaStatus(ctx context.Context, outboxStatus enum.OutboxStatus, sagaStatus ...enum.SagaStatus) []*dto.PaymentOutbox {
	return h.paymentOutboxRepo.FindByTypeAndOutboxStatusAndSagaStatus(ctx, "OrderSaga", outboxStatus, sagaStatus...)
}

func (h *PaymentOutboxHelper) GetPaymentOutboxMessageBySagaIdAndSagaStatus(ctx context.Context, sagaId uuid.UUID, sagaStatus ...enum.SagaStatus) *dto.PaymentOutbox {
	return h.paymentOutboxRepo.FindByTypeAndSagaIdAndSagaStatus(ctx, "OrderSaga", sagaId, sagaStatus...)
}

func (h *PaymentOutboxHelper) SavePaymentOutbox(ctx context.Context, payload *dto.PaymentEventPayload, orderStatus domain.OrderStatus, sagaStatus enum.SagaStatus, outboxStatus enum.OutboxStatus, sagaId uuid.UUID) error {
	payloadStr, err := convertPayload(payload)
	if err != nil {
		return err
	}
	paymentOutbox := &dto.PaymentOutbox{
		ID:           uuid.New(),
		SagaID:       sagaId,
		CreatedAt:    payload.CreatedAt,
		Type:         "OrderSaga",
		Payload:      payloadStr,
		OrderStatus:  orderStatus,
		SagaStatus:   sagaStatus,
		OutboxStatus: outboxStatus,
	}

	if err = h.Save(ctx, paymentOutbox); err != nil {
		return err
	}

	return nil
}

func (h *PaymentOutboxHelper) Save(ctx context.Context, paymentOutbox *dto.PaymentOutbox) error {
	return h.paymentOutboxRepo.Save(ctx, paymentOutbox)
}

func convertPayload(payload *dto.PaymentEventPayload) (string, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("error marshaling to JSON: %s", err)
		return "", err
	}
	return string(jsonData), err
}
