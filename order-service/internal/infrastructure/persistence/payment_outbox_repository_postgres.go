package persistence

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/go-food/order-service/internal/application/outbox"
	"github.com/jpmoraess/go-food/order-service/internal/application/saga"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
	"log"
	"time"
)

type PaymentOutboxEntity struct {
	ID           uuid.UUID
	SagaID       uuid.UUID
	CreatedAt    time.Time
	ProcessedAt  time.Time
	Type         string
	Payload      string
	SagaStatus   saga.SagaStatus
	OrderStatus  domain.OrderStatus
	OutboxStatus outbox.OutboxStatus
	Version      int
}

type PaymentOutboxRepositoryPostgres struct {
	dbpool *pgxpool.Pool
}

func NewPaymentOutboxRepositoryPostgres(dbpool *pgxpool.Pool) *PaymentOutboxRepositoryPostgres {
	return &PaymentOutboxRepositoryPostgres{dbpool: dbpool}
}

func (p *PaymentOutboxRepositoryPostgres) Save(ctx context.Context, paymentOutbox *outbox.PaymentOutbox) error {
	conn, err := p.dbpool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Printf("error starting transaction: %v\n", err)
		return err
	}

	schema := `"order"` // Colocando entre aspas duplas
	table := "payment_outbox"

	query := fmt.Sprintf(`
		INSERT INTO %s.%s (id, saga_id, created_at, processed_at, type, payload, saga_status, order_status, outbox_status, version)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, schema, table)

	_, err = tx.Exec(ctx, query, paymentOutbox.ID, paymentOutbox.SagaID, paymentOutbox.CreatedAt, paymentOutbox.ProcessedAt,
		paymentOutbox.Type, paymentOutbox.Payload, paymentOutbox.SagaStatus, paymentOutbox.OrderStatus, paymentOutbox.OutboxStatus, 1)

	if err != nil {
		err = tx.Rollback(ctx)
		log.Printf("failed to execute insert query: %v\n", err)
		return fmt.Errorf("failed to insert payment outbox entity: %w", err)
	}

	err = tx.Commit(ctx)
	return nil
}

func (p *PaymentOutboxRepositoryPostgres) FindByTypeAndSagaIdAndSagaStatus(ctx context.Context, outboxType string, sagaId uuid.UUID, SagaStatus ...saga.SagaStatus) *outbox.PaymentOutbox {
	return nil
}

func (p *PaymentOutboxRepositoryPostgres) DeleteByTypeAndOutboxStatusAndSagaStatus(ctx context.Context, outboxType string, outboxStatus outbox.OutboxStatus, SagaStatus ...saga.SagaStatus) error {
	return nil
}

func (p *PaymentOutboxRepositoryPostgres) FindByTypeAndOutboxStatusAndSagaStatus(ctx context.Context, outboxType string, outboxStatus outbox.OutboxStatus, SagaStatus ...saga.SagaStatus) []*outbox.PaymentOutbox {
	return nil
}
