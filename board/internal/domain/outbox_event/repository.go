package outbox_event

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type OutboxEventRepository interface {
	Create(ctx context.Context, newEvent NewEvent) error

	WithTx(tx pgx.Tx) OutboxEventRepository
}
