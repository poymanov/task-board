package whoami

import (
	"context"
	"testing"

	jwtMock "github.com/poymanov/codemania-task-board/auth/internal/infrastructure/jwt/mocks"
	"github.com/stretchr/testify/suite"
)

type UseCaseSuite struct {
	suite.Suite

	ctx context.Context // nolint:containedctx

	jwt *jwtMock.JWT

	useCase *UseCase
}

func (s *UseCaseSuite) SetupTest() {
	s.ctx = context.Background()

	s.jwt = jwtMock.NewJWT(s.T())

	s.useCase = NewUseCase(s.jwt)
}

func (s *UseCaseSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}
