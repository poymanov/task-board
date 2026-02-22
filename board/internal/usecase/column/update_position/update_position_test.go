package update_position

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestUpdatePositionError() {
	ErrFailedToCheckBoardExists := errors.New("failed to check if board exists")
	ErrFailedToUpdateColumnPosition := errors.New("failed to update column position")

	tests := []struct {
		name     string
		err      error
		mockFunc func(t *testing.T)
	}{
		{
			name: "Failed to check board exists",
			err:  ErrFailedToCheckBoardExists,
			mockFunc: func(t *testing.T) {
				s.columnRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, ErrFailedToCheckBoardExists).
					Once()
			},
		},
		{
			name: "Column not exists",
			err:  domainBoard.ErrBoardNotExists,
			mockFunc: func(t *testing.T) {
				s.columnRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(false, ErrFailedToCheckBoardExists).
					Once()
			},
		},
		{
			name: "Failed to update column position",
			err:  ErrFailedToUpdateColumnPosition,
			mockFunc: func(t *testing.T) {
				s.columnRepository.
					On("IsExistsById", s.ctx, mock.Anything).
					Return(true, nil).
					Once()
				s.columnRepository.
					On("UpdatePosition", s.ctx, mock.Anything, mock.Anything).
					Return(ErrFailedToUpdateColumnPosition)
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

	s.columnRepository.
		On("IsExistsById", s.ctx, mock.Anything).
		Return(true, nil).
		Once()
	s.columnRepository.
		On("UpdatePosition", s.ctx, mock.Anything, mock.Anything).
		Return(nil)

	err := s.useCase.UpdatePosition(s.ctx, id, UpdatePositionDTO{LeftPosition: gofakeit.Float64(), RightPosition: gofakeit.Float64()})
	s.Require().NoError(err)
}
