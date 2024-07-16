package domain

import "github.com/google/uuid"

type OrderRepository interface {
	Save(order *Order) (*Order, error)
	FindByID(orderID uuid.UUID) (*Order, error)
	FindByTrackingID(trackingID uuid.UUID) (*Order, error)
}
