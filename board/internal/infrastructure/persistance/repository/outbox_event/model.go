package outbox_event

import (
	"time"

	domainOutboxEvent "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event"
)

type OutboxEvent struct {
	Id int `db:"id"`

	EntityType string `db:"entity_type"`

	EntityId string `db:"entity_id"`

	Payload map[string]string `db:"payload"`

	CreatedAt time.Time `db:"created_at"`

	ProcessedAt *time.Time `db:"processed_at"`
}

func ConvertModelToDomain(outboxEvent OutboxEvent) domainOutboxEvent.OutboxEvent {
	return domainOutboxEvent.OutboxEvent{
		Id:          outboxEvent.Id,
		EntityType:  outboxEvent.EntityType,
		EntityId:    outboxEvent.EntityId,
		Payload:     outboxEvent.Payload,
		CreatedAt:   outboxEvent.CreatedAt,
		ProcessedAt: outboxEvent.ProcessedAt,
	}
}
