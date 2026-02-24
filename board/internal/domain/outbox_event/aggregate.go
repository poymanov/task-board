package outbox_event

type NewEvent struct {
	EntityType string

	EntityId int

	Payload map[string]string
}

func NewNewEvent(entityType string, entityId int, payload map[string]string) NewEvent {
	return NewEvent{
		EntityType: entityType,
		EntityId:   entityId,
		Payload:    payload,
	}
}
