package config

import (
	"time"
)

type LoggerConfig interface {
	Level() string
	AppName() string
}

type HttpConfig interface {
	Address() string
	ReadTimeout() time.Duration
}

type HttpMetricsConfig interface {
	Address() string
}

type GrpcClientConfig interface {
	BoardAddress() string
	AuthAddress() string
}

type OtelConfig interface {
	Endpoint() string
	Namespace() string
	InstanceID() string
}
