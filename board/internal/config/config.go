package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/poymanov/codemania-task-board/board/internal/config/env"
)

type Config struct {
	Grpc                GrpcConfig
	Logger              LoggerConfig
	Db                  DbConfig
	Kafka               KafkaConfig
	TaskChangedProducer TaskChangedProducerConfig
	OutboxEvent         OutboxEventConfig
}

func Load(path ...string) (*Config, error) {
	err := godotenv.Load(path...)

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	grpcCfg, err := env.NewGrpcConfig()
	if err != nil {
		return nil, err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return nil, err
	}

	db, err := env.NewDbConfig()
	if err != nil {
		return nil, err
	}

	kafka, err := env.NewKafkaConfig()
	if err != nil {
		return nil, err
	}

	taskChangedProducer, err := env.NewTaskChangedProducerConfig()
	if err != nil {
		return nil, err
	}

	outboxEvent, err := env.NewOutboxEventConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Grpc:                grpcCfg,
		Logger:              loggerCfg,
		Db:                  db,
		Kafka:               kafka,
		TaskChangedProducer: taskChangedProducer,
		OutboxEvent:         outboxEvent,
	}, nil
}
