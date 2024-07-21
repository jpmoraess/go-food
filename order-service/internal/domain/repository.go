package domain

import (
	"context"
	"github.com/google/uuid"
)

type OrderRepository interface {
	Save(ctx context.Context, order *Order) (*Order, error)
	FindByID(ctx context.Context, orderID uuid.UUID) (*Order, error)
	FindByTrackingID(ctx context.Context, trackingID uuid.UUID) (*Order, error)
}
