package create

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestCreateError() {
	ErrFailedToCheckColumnExists := errors.New("failed to check if column exists")
	ErrFailedToCreateTask := errors.New("failed to create task")

	tests := []struct {
		name     string
		err      error
		mockFunc func(t *testing.T)
	}{
		{
			name: "Failed to check column exists",
			err:  ErrFailedToCheckColumnExists,
			mockFunc: func(t *testing.T) {
				s.columnRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, ErrFailedToCheckColumnExists).
					Once()
			},
		},
		{
			name: "Column not exists",
			err:  domainColumn.ErrColumnNotExists,
			mockFunc: func(t *testing.T) {
				s.columnRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, ErrFailedToCheckColumnExists).
					Once()
			},
		},
		{
			name: "Failed to create task",
			err:  ErrFailedToCreateTask,
			mockFunc: func(t *testing.T) {
				s.txManager.
					On("WithTx", mock.Anything, mock.Anything).
					Return(ErrFailedToCreateTask).
					Once()
				s.columnRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(true, nil).
					Once()
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			test.mockFunc(t)
			res, err := s.useCase.Create(s.ctx,
				NewTaskDTO{
					Title:       gofakeit.Name(),
					Description: gofakeit.Name(),
					Assignee:    gofakeit.Name(),
					ColumnId:    int(gofakeit.Int64()),
				},
			)
			assert.Equal(t, 0, res)
			assert.ErrorContains(t, err, err.Error())
		})
	}
}

func (s *UseCaseSuite) TestCreateSuccess() {
	s.columnRepository.
		On("IsExistsById", s.ctx, mock.Anything).
		Return(true, nil).
		Once()
	s.txManager.
		On("WithTx", mock.Anything, mock.Anything).
		Return(nil).
		Once()

	_, err := s.useCase.Create(s.ctx, NewTaskDTO{
		Title:       gofakeit.Name(),
		Description: gofakeit.Name(),
		Assignee:    gofakeit.Name(),
		ColumnId:    int(gofakeit.Int64()),
	})
	s.Require().NoError(err)
}
