package process_event

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type UseCase struct{}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Process(rawData []byte) error {
	var event EventDTO
	if err := json.Unmarshal(rawData, &event); err != nil {
		return err
	}

	log.Info().Any("event", event).Msg("Email отправлен пользователю. Задача обновлена")

	return nil
}
