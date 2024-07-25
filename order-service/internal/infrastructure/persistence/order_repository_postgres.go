package persistence

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type OrderRepositoryPostgres struct {
	dbpool *pgxpool.Pool
}

func NewOrderRepositoryPostgres(dbpool *pgxpool.Pool) *OrderRepositoryPostgres {
	return &OrderRepositoryPostgres{dbpool: dbpool}
}

func (o *OrderRepositoryPostgres) Save(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	conn, err := o.dbpool.Acquire(ctx)
	if err != nil {
		log.Printf("error acquiring connection: %v\n", err)
		return nil, fmt.Errorf("error acquiring connection: %w", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Printf("error starting transaction: %v\n", err)
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	schema := `"order"` // Colocando entre aspas duplas
	table := "orders"

	query := fmt.Sprintf(`
		INSERT INTO %s.%s (id, customer_id, restaurant_id, tracking_id, price, order_status, failure_messages)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, schema, table)

	//query := `INSERT INTO order.orders (id, customer_id, restaurant_id, tracking_id, price, order_status, failure_messages) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = tx.Exec(ctx, query, order.ID(), order.CustomerID(), order.RestaurantID(), order.TrackingID(), order.Price(), order.Status(), strings.Join(order.FailureMessages(), ","))
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			log.Printf("error rolling back transaction: %v\n", rollbackErr)
		}
		log.Printf("failed to execute insert query: %v\n", err)
		return nil, fmt.Errorf("failed to insert order entity: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		log.Printf("error committing transaction: %v\n", err)
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	return order, nil
}

func (o *OrderRepositoryPostgres) FindByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	return nil, nil
}

func (o *OrderRepositoryPostgres) FindByTrackingID(ctx context.Context, trackingID uuid.UUID) (*domain.Order, error) {
	return nil, nil
}
