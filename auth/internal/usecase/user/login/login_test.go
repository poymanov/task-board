package login

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v5"
	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestLoginError() {
	ErrSomeDbError := errors.New("some db error")

	tests := []struct {
		name     string
		err      error
		dto      LoginDTO
		mockFunc func(t *testing.T)
	}{
		{
			name: "User not found",
			err:  domainUser.ErrInvalidCredentials,
			dto:  LoginDTO{},
			mockFunc: func(t *testing.T) {
				s.userRepository.
					On("GetByEmail", s.ctx, mock.Anything).
					Return(domainUser.User{}, pgx.ErrNoRows).
					Once()

				s.userRepository.
					On("IsNoRows", mock.Anything).
					Return(true).
					Once()
			},
		},
		{
			name: "Some db error",
			err:  ErrSomeDbError,
			dto:  LoginDTO{},
			mockFunc: func(t *testing.T) {
				s.userRepository.
					On("GetByEmail", s.ctx, mock.Anything).
					Return(domainUser.User{}, ErrSomeDbError).
					Once()

				s.userRepository.
					On("IsNoRows", mock.Anything).
					Return(false).
					Once()
			},
		},
		{
			name: "Wrong password",
			err:  domainUser.ErrInvalidCredentials,
			dto:  LoginDTO{},
			mockFunc: func(t *testing.T) {
				s.userRepository.
					On("GetByEmail", s.ctx, mock.Anything).
					Return(domainUser.User{Password: "test"}, nil).
					Once()
			},
		},
		{
			name: "Failed to generate access token",
			err:  domainUser.ErrInvalidCredentials,
			dto:  LoginDTO{Password: "123qwe"},
			mockFunc: func(t *testing.T) {
				s.userRepository.
					On("GetByEmail", s.ctx, mock.Anything).
					Return(domainUser.User{Password: "$2a$04$t3ytVVTi5phfsSaQtcFiXuG3kRqBARiHiSlbIk0CPFYd/J3nqz0qi"}, nil).
					Once()

				s.jwt.On("GenerateAccessToken", mock.Anything).
					Return("", errors.New("some jwt error")).
					Once()
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			test.mockFunc(t)
			_, err := s.useCase.Login(s.ctx, test.dto)
			assert.ErrorIs(t, err, test.err)
		})
	}
}

func (s *UseCaseSuite) TestLoginSuccess() {
	accessToken := gofakeit.UUID()

	s.userRepository.
		On("GetByEmail", s.ctx, mock.Anything).
		Return(domainUser.User{Password: "$2a$04$t3ytVVTi5phfsSaQtcFiXuG3kRqBARiHiSlbIk0CPFYd/J3nqz0qi"}, nil).
		Once()

	s.jwt.On("GenerateAccessToken", mock.Anything).
		Return(accessToken, nil).
		Once()

	token, err := s.useCase.Login(s.ctx, LoginDTO{Password: "123qwe"})

	s.Require().NoError(err)
	s.Equal(accessToken, token)
}
