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
	"github.com/poymanov/codemania-task-board/board/internal/config"
	boardRepository "github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/repository/board"
	columnRepository "github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/repository/column"
	outboxEventRepository "github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/repository/outbox_event"
	taskRepository "github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/repository/task"
	"github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/tx_manager"
	transportBoardV1 "github.com/poymanov/codemania-task-board/board/internal/transport/grpc/board/v1/board"
	transportColumnV1 "github.com/poymanov/codemania-task-board/board/internal/transport/grpc/board/v1/column"
	transportTaskV1 "github.com/poymanov/codemania-task-board/board/internal/transport/grpc/board/v1/task"
	boardCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/create"
	boardDeleteUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/delete"
	boardGetAllUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/get_all"
	boardGetBoardUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/get_board"
	columnCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/create"
	columnDeleteUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/delete"
	columnGetAllUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/get_all"
	columnUpdatePositionUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/update_position"
	outboxEventsProcessTasksUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/outbox_event/process_tasks"
	taskCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/create"
	taskGetDeleteUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/delete"
	taskGetAllUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/get_all"
	taskUpdatePositionUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/update_position"
	"github.com/poymanov/codemania-task-board/platform/pkg/grpc/health"
	"github.com/poymanov/codemania-task-board/platform/pkg/logger"
	"github.com/poymanov/codemania-task-board/platform/pkg/migrator"
	"github.com/poymanov/codemania-task-board/platform/pkg/otel"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultInitializationTimeout = time.Second * 10

type App struct {
	closer           []func() error
	listener         net.Listener
	dbConnectionPool *pgxpool.Pool
	config           *config.Config
	txManager        *tx_manager.TxManager
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

	ctxOutboxEvent, cancel := context.WithCancel(ctx)

	defer func() {
		cancel()
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
	a.runOutboxEventProcessTasks(ctxOutboxEvent)

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
		a.initOtel,
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
	a.txManager = tx_manager.NewTxManager(a.dbConnectionPool)

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

func (a *App) initOtel(ctx context.Context) error {
	err := otel.Init(ctx, "board", a.config.Otel.Endpoint(), a.config.Otel.Namespace(), a.config.Otel.InstanceID())
	if err != nil {
		return err
	}

	a.closer = append(a.closer, func() error {
		otel.Close()

		return nil
	})

	return nil
}

func (a *App) InitLogger(_ context.Context) error {
	logger.InitLogger(a.config.Logger.Level(), a.config.Logger.AppName())

	return nil
}

func (a *App) runMigrator() error {
	migration := migrator.NewMigrator(a.dbConnectionPool, a.config.Db.MigrationDirectory())

	if err := migration.Up(); err != nil {
		return err
	}

	return nil
}

func (a *App) runOutboxEventProcessTasks(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(a.config.OutboxEvent.CheckNotProcessedTaskInterval())
		defer ticker.Stop()

		taskChangedProducer := &kafka.Writer{
			Addr:                   kafka.TCP(a.config.Kafka.Brokers()),
			Topic:                  a.config.TaskChangedProducer.Topic(),
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		}

		oer := outboxEventRepository.NewRepository(a.dbConnectionPool)
		oeptuc := outboxEventsProcessTasksUseCase.NewUseCase(oer, a.txManager, taskChangedProducer)

		a.closer = append(a.closer, func() error {
			if err := taskChangedProducer.Close(); err != nil {
				log.Error().Err(err).Msg("failed to close writer")
			}

			return nil
		})

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					return
				default:
				}

				err := oeptuc.Process(ctx, a.config.OutboxEvent.CheckNotProcessedTaskLimit())
				if err != nil {
					log.Error().Err(err).Msg("failed to process outbox events")
				}
			}
		}
	}()
}

func (a *App) runGrpcServer() {
	br := boardRepository.NewRepository(a.dbConnectionPool)
	cr := columnRepository.NewRepository(a.dbConnectionPool)
	tr := taskRepository.NewRepository(a.dbConnectionPool)
	oer := outboxEventRepository.NewRepository(a.dbConnectionPool)

	bcuc := boardCreateUseCase.NewUseCase(br)
	bgauc := boardGetAllUseCase.NewUseCase(br)
	bduc := boardDeleteUseCase.NewUseCase(br)
	bgbuc := boardGetBoardUseCase.NewUseCase(br, cr, tr)

	boardService := transportBoardV1.NewBoardService(bcuc, bgauc, bduc, bgbuc)

	ccuc := columnCreateUseCase.NewUseCase(br, cr)
	cgauc := columnGetAllUseCase.NewUseCase(cr)
	cduc := columnDeleteUseCase.NewUseCase(cr)
	cupuc := columnUpdatePositionUseCase.NewUseCase(cr)

	columnService := transportColumnV1.NewService(ccuc, cgauc, cduc, cupuc)

	tcuc := taskCreateUseCase.NewUseCase(cr, tr, oer, a.txManager)
	tgauc := taskGetAllUseCase.NewUseCase(tr)
	tduc := taskGetDeleteUseCase.NewUseCase(tr, oer, a.txManager)
	tupuc := taskUpdatePositionUseCase.NewUseCase(tr, oer, a.txManager)

	taskService := transportTaskV1.NewService(tcuc, tgauc, tduc, tupuc)

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	boardV1.RegisterBoardServiceServer(s, boardService)
	boardV1.RegisterColumnServiceServer(s, columnService)
	boardV1.RegisterTaskServiceServer(s, taskService)

	health.RegisterService(s)

	reflection.Register(s)

	go func() {
		log.Info().Str("address", a.config.Grpc.Address()).Msg("🚀 gRPC server started")
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
