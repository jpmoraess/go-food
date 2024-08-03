package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/payment-service/internal/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	FindByOrderID(ctx context.Context, orderID uuid.UUID) (*domain.Payment, error)
}
