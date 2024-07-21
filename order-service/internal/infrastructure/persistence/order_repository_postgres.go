package persistence

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
	"log"
)

type OrderEntity struct {
	ID           uuid.UUID
	CustomerID   uuid.UUID
	RestaurantID uuid.UUID
	Price        float64
}

type OrderRepositoryPostgres struct {
	dbpool *pgxpool.Pool
}

func NewOrderRepositoryPostgres(dbpool *pgxpool.Pool) *OrderRepositoryPostgres {
	return &OrderRepositoryPostgres{dbpool: dbpool}
}

func (o *OrderRepositoryPostgres) Save(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	conn, err := o.dbpool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Printf("error starting transaction: %v\n", err)
		return nil, err
	}

	query := ``

	_, err = tx.Exec(ctx, query)
	if err != nil {
		err = tx.Rollback(ctx)
		log.Printf("failed to execute insert query: %v\n", err)
		return nil, fmt.Errorf("failed to insert order entity: %w", err)
	}

	return nil, nil
}

func (o *OrderRepositoryPostgres) FindByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	return nil, nil
}

func (o *OrderRepositoryPostgres) FindByTrackingID(ctx context.Context, trackingID uuid.UUID) (*domain.Order, error) {
	return nil, nil
}
