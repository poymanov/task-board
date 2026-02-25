package outbox_event

import "time"

type OutboxEvent struct {
	Id int

	EntityType string

	EntityId string

	Payload map[string]string

	CreatedAt time.Time

	ProcessedAt *time.Time
}
