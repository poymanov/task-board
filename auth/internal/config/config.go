package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/poymanov/codemania-task-board/auth/internal/config/env"
)

type Config struct {
	Grpc   GrpcConfig
	Logger LoggerConfig
	Db     DbConfig
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

	return &Config{
		Grpc:   grpcCfg,
		Logger: loggerCfg,
		Db:     db,
	}, nil
}
