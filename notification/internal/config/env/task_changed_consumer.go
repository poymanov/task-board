package env

import (
	"github.com/caarlos0/env/v11"
)

type taskChangedConsumerEnvConfig struct {
	TopicName string `env:"TASK_CHANGED_TOPIC_NAME,required"`
	GroupId   string `env:"TASK_CHANGED_CONSUMER_GROUP_ID,required"`
}

type TaskChangedConsumerConfig struct {
	raw taskChangedConsumerEnvConfig
}

func NewTaskChangedConsumerConfig() (*TaskChangedConsumerConfig, error) {
	var raw taskChangedConsumerEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &TaskChangedConsumerConfig{raw: raw}, nil
}

func (cfg *TaskChangedConsumerConfig) Topic() string {
	return cfg.raw.TopicName
}

func (cfg *TaskChangedConsumerConfig) GroupId() string {
	return cfg.raw.GroupId
}
