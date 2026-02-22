package get_all

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestGetAllError() {
	s.taskRepository.
		On("GetAll", s.ctx, mock.Anything, mock.Anything).
		Return([]domainTask.Task{}, errors.New(gofakeit.Word())).Once()

	res, err := s.useCase.GetAll(s.ctx, domainTask.GetAllFilter{}, domainTask.GetAllSort{})
	s.Require().Error(err)
	s.Require().Empty(res)
}

func (s *UseCaseSuite) TestGetAllSuccess() {
	s.taskRepository.
		On("GetAll", s.ctx, mock.Anything, mock.Anything).
		Return([]domainTask.Task{{Id: int(gofakeit.Int64())}}, nil).Once()

	res, err := s.useCase.GetAll(s.ctx, domainTask.GetAllFilter{}, domainTask.GetAllSort{})
	s.Require().NoError(err)
	s.Require().NotEmpty(res)
}
