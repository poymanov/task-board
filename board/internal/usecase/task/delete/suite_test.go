package delete

import (
	"context"
	"testing"

	outboxEventRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event/mocks"
	taskRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/task/mocks"
	txManagerMock "github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/tx_manager/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	taskRepository *taskRepoMock.TaskRepository

	outboxEventRepository *outboxEventRepoMock.OutboxEventRepository

	txManager *txManagerMock.Tx

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.taskRepository = taskRepoMock.NewTaskRepository(s.T())
	s.outboxEventRepository = outboxEventRepoMock.NewOutboxEventRepository(s.T())
	s.txManager = txManagerMock.NewTx(s.T())

	s.useCase = NewUseCase(s.taskRepository, s.outboxEventRepository, s.txManager)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
