package process_tasks

import "time"

type OutboxEventProducerDTO struct {
	Id int `json:"id"`

	EntityType string `json:"entity_type"`

	EntityId string `json:"entity_id"`

	Payload map[string]string `json:"payload"`

	CreatedAt time.Time `json:"created_at"`
}
