package outbox_event

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type OutboxEventRepository interface {
	Create(ctx context.Context, newEvent NewEvent) error

	GetAllNotProcessedByType(ctx context.Context, entityType string, limit int) ([]OutboxEvent, error)

	UpdateEventProcessed(ctx context.Context, id int) error

	WithTx(tx pgx.Tx) OutboxEventRepository
}
