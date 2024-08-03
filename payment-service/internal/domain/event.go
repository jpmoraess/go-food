package domain

import "time"

// PaymentEvent é a interface para eventos de pagamento.
type PaymentEvent interface {
	Payment() *Payment
	CreatedAt() time.Time
	FailureMessages() []string
}

// paymentEvent é uma estrutura base para eventos de pagamento.
type paymentEvent struct {
	payment         *Payment
	createdAt       time.Time
	failureMessages []string
}

// Payment retorna o pagamento associado ao evento.
func (e *paymentEvent) Payment() *Payment {
	return e.payment
}

// CreatedAt retorna a data de criação do evento.
func (e *paymentEvent) CreatedAt() time.Time {
	return e.createdAt
}

// FailureMessages retorna as mensagens de falha do evento.
func (e *paymentEvent) FailureMessages() []string {
	return e.failureMessages
}

// PaymentCompletedEvent representa um evento de pagamento concluído.
type PaymentCompletedEvent struct {
	*paymentEvent
}

// newPaymentCompletedEvent cria um novo evento de pagamento concluído.
func newPaymentCompletedEvent(payment *Payment, createdAt time.Time) PaymentEvent {
	return &PaymentCompletedEvent{
		&paymentEvent{
			payment:         payment,
			createdAt:       createdAt,
			failureMessages: []string{},
		},
	}
}

// PaymentCancelledEvent representa um evento de pagamento cancelado.
type PaymentCancelledEvent struct {
	*paymentEvent
}

func newPaymentCancelledEvent(payment *Payment, createdAt time.Time) PaymentEvent {
	return &PaymentCancelledEvent{
		&paymentEvent{
			payment:         payment,
			createdAt:       createdAt,
			failureMessages: []string{},
		},
	}
}

// PaymentFailedEvent representa um evento de falha de pagamento.
type PaymentFailedEvent struct {
	*paymentEvent
}

func newPaymentFailedEvent(payment *Payment, createdAt time.Time, failureMessages []string) PaymentEvent {
	return &PaymentFailedEvent{
		&paymentEvent{
			payment:         payment,
			createdAt:       createdAt,
			failureMessages: failureMessages,
		},
	}
}
