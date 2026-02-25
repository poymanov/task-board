package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/poymanov/codemania-task-board/notification/internal/config"
	processEventUseCase "github.com/poymanov/codemania-task-board/notification/internal/usecase/process_event"
	"github.com/poymanov/codemania-task-board/platform/pkg/logger"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const defaultInitializationTimeout = time.Second * 10

type App struct {
	closer []func() error
	config *config.Config
}

func newApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.InitDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func Run() (err error) {
	ctx := context.Background()

	a, err := newApp(ctx)
	if err != nil {
		return err
	}

	tccCtx, cancel := context.WithCancel(ctx)

	defer func() {
		cancel()
		ec := a.Close()
		if ec != nil {
			log.Error().Err(ec).Msg("failed to close app")

			err = ec
		}
	}()

	go a.runTaskChangeConsumer(tccCtx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.InitConfig,
		a.InitLogger,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) InitConfig(ctx context.Context) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultInitializationTimeout)
		defer cancel()
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("config initialization timed out")
	default:
	}

	configPath := flag.String("env", ".env", "path to .env file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w, config path: %s", err, *configPath)
	}

	a.config = cfg

	return nil
}

func (a *App) InitLogger(_ context.Context) error {
	logger.InitLogger(a.config.Logger.Level())

	return nil
}

func (a *App) Close() error {
	for i := len(a.closer) - 1; i >= 0; i-- {
		if err := a.closer[i](); err != nil {
			log.Error().Err(err).Msg("failed to close application component")
		}
	}

	return nil
}

func (a *App) runTaskChangeConsumer(ctx context.Context) {
	peuc := processEventUseCase.NewUseCase()

	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           a.config.Kafka.Brokers(),
		Topic:             a.config.TaskChangedConsumer.Topic(),
		GroupID:           a.config.TaskChangedConsumer.GroupId(),
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
		RebalanceTimeout:  60 * time.Second,
	})

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close consumer")
		}
	}()

	for {
		m, err := consumer.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				log.Info().Err(err).Msg("consumer stopped by context")
				return
			}

			log.Error().Err(err).Msg("failed to read message")

			continue
		}

		err = peuc.Process(m.Value)
		if err != nil {
			log.Error().Err(err).Msg("failed to process message")
		}
	}
}
