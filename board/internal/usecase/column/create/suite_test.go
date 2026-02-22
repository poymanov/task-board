package create

import (
	"context"
	"testing"

	boardRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/board/mocks"
	columnRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/column/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	boardRepository *boardRepoMock.BoardRepository

	columnRepository *columnRepoMock.ColumnRepository

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.boardRepository = boardRepoMock.NewBoardRepository(s.T())
	s.columnRepository = columnRepoMock.NewColumnRepository(s.T())

	s.useCase = NewUseCase(s.boardRepository, s.columnRepository)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
