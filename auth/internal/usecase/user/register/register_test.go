package register

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestRegisterUserError() {
	ErrSomeDbError := errors.New("some db error")

	tests := []struct {
		name     string
		err      error
		mockFunc func(t *testing.T)
	}{
		{
			name: "User already exists",
			err:  domainUser.ErrUserAlreadyExists,
			mockFunc: func(t *testing.T) {
				ErrUniqueViolation := pgconn.PgError{Code: "23505"}

				s.userRepository.
					On("Create", s.ctx, mock.Anything).
					Return(&ErrUniqueViolation).
					Once()

				s.userRepository.
					On("IsDuplicateKey", mock.Anything).
					Return(true).
					Once()
			},
		},
		{
			name: "Some db error",
			err:  ErrSomeDbError,
			mockFunc: func(t *testing.T) {
				s.userRepository.
					On("Create", s.ctx, mock.Anything).
					Return(ErrSomeDbError).
					Once()

				s.userRepository.
					On("IsDuplicateKey", mock.Anything).
					Return(false).
					Once()
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			test.mockFunc(t)
			err := s.useCase.Register(s.ctx, RegisterUserDTO{})
			assert.ErrorIs(t, err, test.err)
		})
	}
}

func (s *UseCaseSuite) TestRegisterUserSuccess() {
	s.userRepository.
		On("Create", s.ctx, mock.Anything).
		Return(nil).
		Once()

	err := s.useCase.Register(s.ctx, RegisterUserDTO{})

	s.Require().NoError(err)
}
