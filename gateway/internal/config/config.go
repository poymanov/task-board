package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/poymanov/codemania-task-board/gateway/internal/config/env"
)

type Config struct {
	Logger      LoggerConfig
	Http        HttpConfig
	HttpMetrics HttpMetricsConfig
	GrpcClient  GrpcClientConfig
	Otel        OtelConfig
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

	httpCfg, err := env.NewHttpConfig()
	if err != nil {
		return nil, err
	}

	httpMetricsCfg, err := env.NewHttpMetricsConfig()
	if err != nil {
		return nil, err
	}

	grpcClient, err := env.NewGrpcClient()
	if err != nil {
		return nil, err
	}

	otel, err := env.NewOtelConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Logger:      loggerCfg,
		Http:        httpCfg,
		HttpMetrics: httpMetricsCfg,
		GrpcClient:  grpcClient,
		Otel:        otel,
	}, nil
}
