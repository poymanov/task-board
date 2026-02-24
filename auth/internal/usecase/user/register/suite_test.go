package register

import (
	"context"
	"testing"

	repoMock "github.com/poymanov/codemania-task-board/auth/internal/domain/user/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	userRepository *repoMock.UserRepository

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.userRepository = repoMock.NewUserRepository(s.T())

	s.useCase = NewUseCase(s.userRepository)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
