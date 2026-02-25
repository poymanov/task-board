package env

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type outboxEventEnvConfig struct {
	CheckNotProcessedTaskInterval time.Duration `env:"BOARD_CHECK_NOT_PROCESSED_TASKS_INTERVAL,required"`
	CheckNotProcessedTaskLimit    int           `env:"BOARD_CHECK_PROCESSED_TASKS_INTERVAL_LIMIT,required"`
}

type OutboxEventConfig struct {
	raw outboxEventEnvConfig
}

func NewOutboxEventConfig() (*OutboxEventConfig, error) {
	var raw outboxEventEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &OutboxEventConfig{raw: raw}, nil
}

func (o *OutboxEventConfig) CheckNotProcessedTaskInterval() time.Duration {
	return o.raw.CheckNotProcessedTaskInterval
}

func (o *OutboxEventConfig) CheckNotProcessedTaskLimit() int {
	return o.raw.CheckNotProcessedTaskLimit
}
