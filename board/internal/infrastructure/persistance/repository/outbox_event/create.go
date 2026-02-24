package outbox_event

import (
	"context"
	"fmt"

	domainOutboxEvent "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event"
)

func (r *Repository) Create(ctx context.Context, newEvent domainOutboxEvent.NewEvent) error {
	_, err := r.pool.Exec(
		ctx,
		`INSERT INTO outbox_events (entity_type, entity_id, payload) VALUES ($1,$2,$3)`,
		newEvent.EntityType, newEvent.EntityId, newEvent.Payload,
	)
	if err != nil {
		return fmt.Errorf("failed to create outbox event: %w", err)
	}

	return nil
}
