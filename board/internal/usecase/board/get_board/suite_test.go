package get_board

import (
	"context"
	"testing"

	boardRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/board/mocks"
	columnRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/column/mocks"
	taskRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/task/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	boardRepository *boardRepoMock.BoardRepository

	columnRepository *columnRepoMock.ColumnRepository

	taskRepository *taskRepoMock.TaskRepository

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.boardRepository = boardRepoMock.NewBoardRepository(s.T())
	s.columnRepository = columnRepoMock.NewColumnRepository(s.T())
	s.taskRepository = taskRepoMock.NewTaskRepository(s.T())

	s.useCase = NewUseCase(s.boardRepository, s.columnRepository, s.taskRepository)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
