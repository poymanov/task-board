package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/poymanov/codemania-task-board/auth/internal/config"
	"github.com/poymanov/codemania-task-board/auth/internal/infrastructure/jwt"
	userRepository "github.com/poymanov/codemania-task-board/auth/internal/infrastructure/persistance/repository/user"
	transportUserV1 "github.com/poymanov/codemania-task-board/auth/internal/transport/grpc/auth/v1/user"
	loginUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/login"
	registerUserUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/register"
	whoamiUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/whoami"
	"github.com/poymanov/codemania-task-board/platform/pkg/grpc/health"
	"github.com/poymanov/codemania-task-board/platform/pkg/logger"
	"github.com/poymanov/codemania-task-board/platform/pkg/migrator"
	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultInitializationTimeout = time.Second * 10

type App struct {
	closer           []func() error
	listener         net.Listener
	dbConnectionPool *pgxpool.Pool
	config           *config.Config
}

func newApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.InitDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func Run() error {
	ctx := context.Background()

	a, err := newApp(ctx)
	if err != nil {
		return err
	}

	defer func() {
		ec := a.Close()
		if ec != nil {
			log.Error().Err(ec).Msg("failed to close app")
			return
		}
	}()

	err = a.runMigrator()
	if err != nil {
		return err
	}

	a.runGrpcServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func (a *App) InitConfig(ctx context.Context) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultInitializationTimeout)
		defer cancel()
	}

	chDone := make(chan struct{})

	var configErr error

	go func() {
		configPath := flag.String("env", ".env", "path to .env file")
		flag.Parse()

		cfg, err := config.Load(*configPath)

		if err != nil {
			configErr = fmt.Errorf("failed to load config: %w, config path: %s", err, *configPath)
		} else {
			a.config = cfg
		}

		chDone <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		configErr = fmt.Errorf("config initialization timed out")
	case <-chDone:
	}

	if configErr != nil {
		return configErr
	}

	return nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.InitConfig,
		a.InitLogger,
		a.InitDB,
		a.initListener,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) InitDB(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, a.config.Db.Uri())
	if err != nil {
		log.Error().Err(err).Msg("db failed connect")
		return err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Error().Err(err).Msg("db not available")
		return err

	}

	a.dbConnectionPool = pool

	a.closer = append(a.closer, func() error {
		pool.Close()

		return nil
	})

	return nil
}

func (a *App) initListener(_ context.Context) error {
	list, err := net.Listen("tcp", a.config.Grpc.Address())
	if err != nil {
		log.Error().Err(err).Msg("failed to start listener")
		return err
	}

	a.listener = list

	a.closer = append(a.closer, func() error {
		lerr := list.Close()

		if lerr != nil && !errors.Is(lerr, net.ErrClosed) {
			log.Error().Err(err).Msg("failed to close listener")

			return lerr
		}

		return nil
	})

	return nil
}

func (a *App) InitLogger(_ context.Context) error {
	logger.InitLogger(a.config.Logger.Level())

	return nil
}

func (a *App) runMigrator() error {
	migration := migrator.NewMigrator(a.dbConnectionPool, a.config.Db.MigrationDirectory())

	if err := migration.Up(); err != nil {
		return err
	}

	return nil
}

func (a *App) runGrpcServer() {
	ur := userRepository.NewRepository(a.dbConnectionPool)
	js := jwt.NewJWTService(a.config.JWT.AccessTokenTTL(), a.config.JWT.AccessTokenSecret())

	ruuc := registerUserUseCase.NewUseCase(ur)
	luc := loginUseCase.NewUseCase(ur, js)
	wuc := whoamiUseCase.NewUseCase(js)

	userService := transportUserV1.NewService(ruuc, luc, wuc)

	s := grpc.NewServer()

	authV1.RegisterUserServiceServer(s, userService)

	health.RegisterService(s)

	reflection.Register(s)

	go func() {
		log.Info().Msg(fmt.Sprintf("🚀 gRPC server listening on %s\n", a.config.Grpc.Address()))
		err := s.Serve(a.listener)
		if err != nil {
			log.Error().Err(err).Msg("failed to serve grpc server")
			return
		}
	}()

	a.closer = append(a.closer, func() error {
		s.GracefulStop()

		return nil
	})
}

func (a *App) Close() error {
	for i := len(a.closer) - 1; i >= 0; i-- {
		if err := a.closer[i](); err != nil {
			log.Error().Err(err).Msg("failed to close application component")
		}
	}

	return nil
}
