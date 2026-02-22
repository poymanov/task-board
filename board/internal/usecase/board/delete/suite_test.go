package delete

import (
	"context"
	"testing"

	repoMock "github.com/poymanov/codemania-task-board/board/internal/domain/board/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	boardRepository *repoMock.BoardRepository

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.boardRepository = repoMock.NewBoardRepository(s.T())

	s.useCase = NewUseCase(s.boardRepository)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
