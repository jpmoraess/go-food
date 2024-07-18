package outbox

import (
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

func (h *PaymentOutboxHelper) GetPaymentOutboxByOutboxStatusAndSagaStatus(outboxStatus OutboxStatus, sagaStatus ...saga.SagaStatus) []*PaymentOutbox {
	return h.paymentOutboxRepo.FindByTypeAndOutboxStatusAndSagaStatus("OrderSaga", outboxStatus, sagaStatus...)
}

func (h *PaymentOutboxHelper) GetPaymentOutboxMessageBySagaIdAndSagaStatus(sagaId uuid.UUID, sagaStatus ...saga.SagaStatus) *PaymentOutbox {
	return h.paymentOutboxRepo.FindByTypeAndSagaIdAndSagaStatus("OrderSaga", sagaId, sagaStatus...)
}

func (h *PaymentOutboxHelper) SavePaymentOutbox(payload *PaymentEventPayload, orderStatus domain.OrderStatus, sagaStatus saga.SagaStatus, outboxStatus OutboxStatus, sagaId uuid.UUID) error {
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

	if err = h.save(paymentOutbox); err != nil {
		return err
	}

	return nil
}

func (h *PaymentOutboxHelper) save(paymentOutbox *PaymentOutbox) error {
	return h.paymentOutboxRepo.Save(paymentOutbox)
}

func convertPayload(payload *PaymentEventPayload) (string, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("error marshaling to JSON: %s", err)
		return "", err
	}
	return string(jsonData), err
}
