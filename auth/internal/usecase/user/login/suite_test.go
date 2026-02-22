package login

import (
	"context"
	"testing"

	repoMock "github.com/poymanov/codemania-task-board/auth/internal/domain/user/mocks"
	jwtMock "github.com/poymanov/codemania-task-board/auth/internal/infrastructure/jwt/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	userRepository *repoMock.UserRepository

	jwt *jwtMock.JWT

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.userRepository = repoMock.NewUserRepository(s.T())
	s.jwt = jwtMock.NewJWT(s.T())

	s.useCase = NewUseCase(s.userRepository, s.jwt)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
