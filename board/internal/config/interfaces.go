package config

import "time"

type GrpcConfig interface {
	Address() string
}

type LoggerConfig interface {
	Level() string
	AppName() string
}

type DbConfig interface {
	Uri() string
	MigrationDirectory() string
}

type KafkaConfig interface {
	Brokers() string
}

type TaskChangedProducerConfig interface {
	Topic() string
}

type OutboxEventConfig interface {
	CheckNotProcessedTaskInterval() time.Duration
	CheckNotProcessedTaskLimit() int
}

type OtelConfig interface {
	Endpoint() string
	Namespace() string
	InstanceID() string
}
