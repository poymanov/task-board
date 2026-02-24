package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type dbEnvConfig struct {
	Host               string `env:"POSTGRES_AUTH_HOST,required"`
	Port               string `env:"POSTGRES_AUTH_PORT,required"`
	User               string `env:"POSTGRES_AUTH_USER,required"`
	Password           string `env:"POSTGRES_AUTH_PASSWORD,required"`
	DBName             string `env:"POSTGRES_AUTH_DB,required"`
	SSLMode            string `env:"POSTGRES_AUTH_SSL_MODE,required"`
	MigrationDirectory string `env:"MIGRATION_AUTH_DIRECTORY,required"`
}

type DbConfig struct {
	raw dbEnvConfig
}

func NewDbConfig() (*DbConfig, error) {
	var raw dbEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &DbConfig{raw: raw}, nil
}

func (cfg *DbConfig) Uri() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.raw.User,
		cfg.raw.Password,
		cfg.raw.Host,
		cfg.raw.Port,
		cfg.raw.DBName,
		cfg.raw.SSLMode)
}

func (cfg *DbConfig) MigrationDirectory() string {
	return cfg.raw.MigrationDirectory
}
