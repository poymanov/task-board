package config

import "time"

type GrpcConfig interface {
	Address() string
}

type LoggerConfig interface {
	Level() string
}

type DbConfig interface {
	Uri() string
	MigrationDirectory() string
}

type JWTConfig interface {
	AccessTokenTTL() time.Duration
	AccessTokenSecret() string
}
