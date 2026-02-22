package update_position

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestUpdatePositionError() {
	ErrFailedToCheckTaskExists := errors.New("failed to check if task exists")
	ErrFailedToUpdateTaskPosition := errors.New("failed to update task position")

	tests := []struct {
		name     string
		err      error
		mockFunc func(t *testing.T)
	}{
		{
			name: "Failed to check task exists",
			err:  ErrFailedToCheckTaskExists,
			mockFunc: func(t *testing.T) {
				s.taskRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, ErrFailedToCheckTaskExists).
					Once()
			},
		},
		{
			name: "Task not exists",
			err:  domainTask.ErrTaskNotExists,
			mockFunc: func(t *testing.T) {
				s.taskRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, domainTask.ErrTaskNotExists).
					Once()
			},
		},
		{
			name: "Failed to update task position",
			err:  ErrFailedToUpdateTaskPosition,
			mockFunc: func(t *testing.T) {
				s.taskRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(true, nil).
					Once()
				s.taskRepository.
					On("UpdatePosition", s.ctx, mock.Anything, mock.Anything).
					Return(ErrFailedToUpdateTaskPosition)
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			id := int(gofakeit.Int64())
			test.mockFunc(t)
			err := s.useCase.UpdatePosition(s.ctx, id, UpdatePositionDTO{LeftPosition: gofakeit.Float64(), RightPosition: gofakeit.Float64()})
			assert.ErrorContains(t, err, err.Error())
		})
	}
}

func (s *UseCaseSuite) TestUpdatePositionSuccess() {
	id := int(gofakeit.Int64())

	s.taskRepository.
		On("IsExistsById", s.ctx, mock.Anything).
		Return(true, nil).
		Once()
	s.taskRepository.
		On("UpdatePosition", s.ctx, mock.Anything, mock.Anything).
		Return(nil)

	err := s.useCase.UpdatePosition(s.ctx, id, UpdatePositionDTO{LeftPosition: gofakeit.Float64(), RightPosition: gofakeit.Float64()})
	s.Require().NoError(err)
}
