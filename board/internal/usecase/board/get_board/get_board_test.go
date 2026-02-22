package get_board

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	domainCommon "github.com/poymanov/codemania-task-board/board/internal/domain/common"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestGetBoardError() {
	ErrFailedToGetBoard := errors.New("failed to get board")
	ErrFailedToGetColumns := errors.New("failed to get columns")
	ErrFailedToGetTasks := errors.New("failed to get tasks")

	tests := []struct {
		name     string
		err      error
		mockFunc func(t *testing.T)
	}{
		{
			name: "Failed to get board",
			err:  ErrFailedToGetBoard,
			mockFunc: func(t *testing.T) {
				s.boardRepository.
					On("GetById", s.ctx, mock.Anything).
					Return(domainBoard.Board{}, ErrFailedToGetBoard).
					Once()
			},
		},
		{
			name: "Failed to get columns",
			err:  ErrFailedToGetColumns,
			mockFunc: func(t *testing.T) {
				s.boardRepository.
					On("GetById", s.ctx, mock.Anything).
					Return(domainBoard.Board{}, nil).
					Once()

				s.columnRepository.
					On("GetAll", s.ctx, mock.Anything, mock.Anything).
					Return([]domainColumn.Column{}, ErrFailedToGetColumns).
					Once()
			},
		},
		{
			name: "Failed to get tasks",
			err:  ErrFailedToGetTasks,
			mockFunc: func(t *testing.T) {
				s.boardRepository.
					On("GetById", s.ctx, mock.Anything).
					Return(domainBoard.Board{}, nil).
					Once()

				s.columnRepository.
					On("GetAll", s.ctx, mock.Anything, mock.Anything).
					Return([]domainColumn.Column{}, nil).
					Once()

				s.taskRepository.
					On("GetAll", s.ctx, mock.Anything, mock.Anything).
					Return([]domainTask.Task{}, ErrFailedToGetTasks).
					Once()
			},
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			test.mockFunc(t)
			res, err := s.useCase.GetBoard(s.ctx, 1)
			assert.Equal(t, domainCommon.Board{}, res)
			assert.ErrorContains(t, err, err.Error())
		})
	}
}

func (s *UseCaseSuite) TestGetBoardSuccess() {
	s.boardRepository.
		On("GetById", s.ctx, mock.Anything).
		Return(domainBoard.Board{}, nil).
		Once()

	s.columnRepository.
		On("GetAll", s.ctx, mock.Anything, mock.Anything).
		Return([]domainColumn.Column{}, nil).
		Once()

	s.taskRepository.
		On("GetAll", s.ctx, mock.Anything, mock.Anything).
		Return([]domainTask.Task{}, nil).
		Once()

	res, err := s.useCase.GetBoard(s.ctx, int(gofakeit.Int64()))

	s.Require().NoError(err)
	s.Require().NotEmpty(res)
}
