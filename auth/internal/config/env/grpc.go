package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type grpcEnvConfig struct {
	Host string `env:"AUTH_GRPC_HOST,required"`
	Port string `env:"AUTH_GRPC_PORT,required"`
}

type GrpcConfig struct {
	raw grpcEnvConfig
}

func NewGrpcConfig() (*GrpcConfig, error) {
	var raw grpcEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &GrpcConfig{raw: raw}, nil
}

func (cfg *GrpcConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}
