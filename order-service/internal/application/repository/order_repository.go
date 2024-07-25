package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, order *domain.Order) (*domain.Order, error)
	FindByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error)
	FindByTrackingID(ctx context.Context, trackingID uuid.UUID) (*domain.Order, error)
}
