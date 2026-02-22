package delete

import (
	"context"
	"testing"

	taskRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/task/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	taskRepository *taskRepoMock.TaskRepository

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.taskRepository = taskRepoMock.NewTaskRepository(s.T())

	s.useCase = NewUseCase(s.taskRepository)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
