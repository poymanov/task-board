package whoami

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestWhoamiError() {
	accessToken := gofakeit.UUID()

	ErrJwt := errors.New("some jwt error")

	s.jwt.
		On("ValidateAccessToken", mock.Anything).
		Return(domainUser.AuthClaims{}, ErrJwt).Once()

	res, err := s.useCase.Whoami(accessToken)

	s.Require().Error(err)
	s.Require().Empty(res)
}

func (s *UseCaseSuite) TestWhoamiSuccess() {
	accessToken := gofakeit.UUID()

	s.jwt.
		On("ValidateAccessToken", mock.Anything).
		Return(domainUser.AuthClaims{Email: gofakeit.Email()}, nil).
		Once()

	res, err := s.useCase.Whoami(accessToken)

	s.Require().NoError(err)
	s.Require().NotEmpty(res)
}
