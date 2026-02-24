package main

import (
	"github.com/poymanov/codemania-task-board/auth/internal/app"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Error().Err(err).Msg("failed to run app")
	}
}
