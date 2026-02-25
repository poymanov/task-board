package env

import "github.com/caarlos0/env/v11"

type taskChangedProducerEnvConfig struct {
	TopicName string `env:"TASK_CHANGED_TOPIC_NAME,required"`
}

type TaskChangedProducerConfig struct {
	raw taskChangedProducerEnvConfig
}

func NewTaskChangedProducerConfig() (*TaskChangedProducerConfig, error) {
	var raw taskChangedProducerEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &TaskChangedProducerConfig{raw: raw}, nil
}

func (cfg *TaskChangedProducerConfig) Topic() string {
	return cfg.raw.TopicName
}
