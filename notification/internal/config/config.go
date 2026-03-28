package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/poymanov/codemania-task-board/notification/internal/config/env"
)

type Config struct {
	Logger              LoggerConfig
	Kafka               KafkaConfig
	TaskChangedConsumer TaskChangedConsumerConfig
}

func Load(path ...string) (*Config, error) {
	err := godotenv.Load(path...)

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return nil, err
	}

	kafka, err := env.NewKafkaConfig()
	if err != nil {
		return nil, err
	}

	taskChangedConsumer, err := env.NewTaskChangedConsumerConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Logger:              loggerCfg,
		Kafka:               kafka,
		TaskChangedConsumer: taskChangedConsumer,
	}, nil
}
