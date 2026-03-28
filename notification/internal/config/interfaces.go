package config

type LoggerConfig interface {
	Level() string
}

type KafkaConfig interface {
	Brokers() []string
}

type TaskChangedConsumerConfig interface {
	Topic() string
	GroupId() string
}
