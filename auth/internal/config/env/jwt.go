package env

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type jwtEnvConfig struct {
	AccessTokenTTL    time.Duration `env:"AUTH_JWT_ACCESS_TOKEN_TTL,required"`
	AccessTokenSecret string        `env:"AUTH_JWT_ACCESS_TOKEN_SECRET,required"`
}

type JWTConfig struct {
	raw jwtEnvConfig
}

func NewJWTConfig() (*JWTConfig, error) {
	var raw jwtEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &JWTConfig{raw: raw}, nil
}

func (cfg JWTConfig) AccessTokenTTL() time.Duration {
	return cfg.raw.AccessTokenTTL
}

func (cfg JWTConfig) AccessTokenSecret() string {
	return cfg.raw.AccessTokenSecret
}
