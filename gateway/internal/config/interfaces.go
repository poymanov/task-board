package config

import (
	"time"
)

type LoggerConfig interface {
	Level() string
}

type HttpConfig interface {
	Address() string
	ReadTimeout() time.Duration
}

type GrpcClientConfig interface {
	BoardAddress() string
	AuthAddress() string
}
