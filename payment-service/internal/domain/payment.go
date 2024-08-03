package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	paymentID     uuid.UUID
	orderID       uuid.UUID
	customerID    uuid.UUID
	price         float64
	paymentStatus string
	createdAt     time.Time
}

func NewPayment(orderID, customerID uuid.UUID, price float64) *Payment {
	return &Payment{
		paymentID:     uuid.New(),
		orderID:       orderID,
		customerID:    customerID,
		price:         price,
		paymentStatus: "STARTED",
		createdAt:     time.Now(),
	}
}

func (p *Payment) validatePayment(failureMessages *[]string) {
	if !(p.price > 0) {
		*failureMessages = append(*failureMessages, "Total price must be greater than zero.")
	}
}

func (p *Payment) updateStatus(status string) {
	p.paymentStatus = status
}
