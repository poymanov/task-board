package get_all

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestGetAllError() {
	ownerId := int(gofakeit.Int64())

	s.boardRepository.
		On("GetAll", s.ctx, mock.Anything).
		Return([]domainBoard.Board{}, errors.New(gofakeit.Word())).Once()

	res, err := s.useCase.GetAll(s.ctx, domainBoard.GetAllFilter{OwnerId: ownerId})
	s.Require().Error(err)
	s.Require().Empty(res)
}

func (s *UseCaseSuite) TestGetAllSuccess() {
	ownerId := int(gofakeit.Int64())
	s.boardRepository.On("GetAll", s.ctx, mock.Anything).Return([]domainBoard.Board{{OwnerId: ownerId}}, nil).Once()

	res, err := s.useCase.GetAll(s.ctx, domainBoard.GetAllFilter{OwnerId: ownerId})
	s.Require().NoError(err)
	s.Require().NotEmpty(res)
}
