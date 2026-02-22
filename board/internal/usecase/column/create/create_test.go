package create

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestCreateError() {
	ErrFailedToCheckBoardExists := errors.New("failed to check if board exists")
	ErrFailedToCreateColumn := errors.New("failed to create column")

	tests := []struct {
		name     string
		err      error
		mockFunc func(t *testing.T)
	}{
		{
			name: "Failed to check board exists",
			err:  ErrFailedToCheckBoardExists,
			mockFunc: func(t *testing.T) {
				s.boardRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, ErrFailedToCheckBoardExists).
					Once()
			},
		},
		{
			name: "Board not exists",
			err:  domainBoard.ErrBoardNotExists,
			mockFunc: func(t *testing.T) {
				s.boardRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, ErrFailedToCheckBoardExists).
					Once()
			},
		},
		{
			name: "Failed to create column",
			err:  ErrFailedToCreateColumn,
			mockFunc: func(t *testing.T) {
				s.boardRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(true, nil).
					Once()
				s.columnRepository.
					On("Create", mock.Anything, mock.Anything).
					Return(0, ErrFailedToCreateColumn)
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			test.mockFunc(t)
			res, err := s.useCase.Create(s.ctx, NewColumnDTO{Name: gofakeit.Name(), BoardID: int(gofakeit.Int64())})
			assert.Equal(t, 0, res)
			assert.ErrorContains(t, err, err.Error())
		})
	}
}

func (s *UseCaseSuite) TestCreateSuccess() {
	columnId := int(gofakeit.Int64())
	s.boardRepository.
		On("IsExistsById", s.ctx, mock.Anything).
		Return(true, nil).
		Once()
	s.columnRepository.
		On("Create", mock.Anything, mock.Anything).
		Return(columnId, nil)

	res, err := s.useCase.Create(s.ctx, NewColumnDTO{Name: gofakeit.Name(), BoardID: int(gofakeit.Int64())})
	s.Require().NoError(err)
	s.Require().Equal(columnId, res)
}
