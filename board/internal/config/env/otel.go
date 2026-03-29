package env

import "github.com/caarlos0/env/v11"

type otelEnvConfig struct {
	Endpoint   string `env:"BOARD_OTEL_ENDPOINT,required"`
	Namespace  string `env:"BOARD_OTEL_NAMESPACE,required"`
	InstanceID string `env:"BOARD_OTEL_INSTANCE_ID,required"`
}

type OtelConfig struct {
	raw otelEnvConfig
}

func NewOtelConfig() (*OtelConfig, error) {
	var raw otelEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &OtelConfig{raw: raw}, nil
}

func (cfg *OtelConfig) Endpoint() string {
	return cfg.raw.Endpoint
}

func (cfg *OtelConfig) Namespace() string {
	return cfg.raw.Namespace
}

func (cfg *OtelConfig) InstanceID() string {
	return cfg.raw.InstanceID
}
