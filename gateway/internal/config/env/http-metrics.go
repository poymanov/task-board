package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type httpMetricsEnvConfig struct {
	Port string `env:"GATEWAY_HTTP_METRICS_PORT,required"`
}

type HttpMetricsConfig struct {
	raw httpMetricsEnvConfig
}

func NewHttpMetricsConfig() (*HttpMetricsConfig, error) {
	var raw httpMetricsEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &HttpMetricsConfig{raw: raw}, nil
}

func (cfg *HttpMetricsConfig) Address() string {
	return net.JoinHostPort("", cfg.raw.Port)
}
