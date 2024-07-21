package outbox

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
	"log"
)

type PaymentOutboxHelper struct {
	paymentOutboxRepo PaymentOutboxRepository
}

func NewPaymentOutboxHelper(paymentOutboxRepo PaymentOutboxRepository) *PaymentOutboxHelper {
	return &PaymentOutboxHelper{paymentOutboxRepo: paymentOutboxRepo}
}

func (h *PaymentOutboxHelper) GetPaymentOutboxByOutboxStatusAndSagaStatus(ctx context.Context, outboxStatus OutboxStatus, sagaStatus ...saga.SagaStatus) []*PaymentOutbox {
	return h.paymentOutboxRepo.FindByTypeAndOutboxStatusAndSagaStatus(ctx, "OrderSaga", outboxStatus, sagaStatus...)
}

func (h *PaymentOutboxHelper) GetPaymentOutboxMessageBySagaIdAndSagaStatus(ctx context.Context, sagaId uuid.UUID, sagaStatus ...saga.SagaStatus) *PaymentOutbox {
	return h.paymentOutboxRepo.FindByTypeAndSagaIdAndSagaStatus(ctx, "OrderSaga", sagaId, sagaStatus...)
}

func (h *PaymentOutboxHelper) SavePaymentOutbox(ctx context.Context, payload *PaymentEventPayload, orderStatus domain.OrderStatus, sagaStatus saga.SagaStatus, outboxStatus OutboxStatus, sagaId uuid.UUID) error {
	id, _ := uuid.NewRandom()
	payloadStr, err := convertPayload(payload)
	if err != nil {
		return err
	}
	paymentOutbox := &PaymentOutbox{
		ID:           id,
		SagaID:       sagaId,
		CreatedAt:    payload.CreatedAt,
		Type:         "OrderSaga",
		Payload:      payloadStr,
		OrderStatus:  orderStatus,
		SagaStatus:   sagaStatus,
		OutboxStatus: outboxStatus,
	}

	if err = h.save(ctx, paymentOutbox); err != nil {
		return err
	}

	return nil
}

func (h *PaymentOutboxHelper) save(ctx context.Context, paymentOutbox *PaymentOutbox) error {
	return h.paymentOutboxRepo.Save(ctx, paymentOutbox)
}

func convertPayload(payload *PaymentEventPayload) (string, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("error marshaling to JSON: %s", err)
		return "", err
	}
	return string(jsonData), err
}
