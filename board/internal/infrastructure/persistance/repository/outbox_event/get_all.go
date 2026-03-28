package outbox_event

import (
	"context"

	"github.com/jackc/pgx/v5"
	domainOutboxEvent "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event"
)

func (r *Repository) GetAllNotProcessedByType(ctx context.Context, entityType string, limit int) ([]domainOutboxEvent.OutboxEvent, error) {
	rows, err := r.pool.Query(
		ctx,
		`SELECT * FROM outbox_events WHERE entity_type = $1 AND processed_at IS NULL LIMIT $2`,
		entityType, limit,
	)
	if err != nil {
		return []domainOutboxEvent.OutboxEvent{}, err
	}

	defer rows.Close()

	models, err := pgx.CollectRows(rows, pgx.RowToStructByName[OutboxEvent])
	if err != nil {
		return []domainOutboxEvent.OutboxEvent{}, err
	}

	outboxEvents := make([]domainOutboxEvent.OutboxEvent, 0, len(models))

	for _, model := range models {
		outboxEvents = append(outboxEvents, ConvertModelToDomain(model))
	}

	return outboxEvents, nil
}
