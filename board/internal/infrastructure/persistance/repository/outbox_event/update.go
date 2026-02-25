package outbox_event

import (
	"context"
	"fmt"
)

func (r *Repository) UpdateEventProcessed(ctx context.Context, id int) error {
	_, err := r.pool.Exec(
		ctx,
		`UPDATE outbox_events SET processed_at = now() WHERE id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to update event processed: %w", err)
	}

	return nil
}
