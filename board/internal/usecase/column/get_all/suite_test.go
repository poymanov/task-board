package get_all

import (
	"context"
	"testing"

	columnRepoMock "github.com/poymanov/codemania-task-board/board/internal/domain/column/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	columnRepository *columnRepoMock.ColumnRepository

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.columnRepository = columnRepoMock.NewColumnRepository(s.T())

	s.useCase = NewUseCase(s.columnRepository)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
