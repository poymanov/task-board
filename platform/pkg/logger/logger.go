package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(lvl, appName string) {
	level, err := zerolog.ParseLevel(lvl)
	if err != nil {
		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)

	log.Logger = log.With().
		Str("app_name", appName).
		Logger()
}
