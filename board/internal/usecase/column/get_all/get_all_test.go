package get_all

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestGetAllError() {
	s.columnRepository.
		On("GetAll", s.ctx, mock.Anything, mock.Anything).
		Return([]domainColumn.Column{}, errors.New(gofakeit.Word())).Once()

	res, err := s.useCase.GetAll(s.ctx, domainColumn.GetAllFilter{}, domainColumn.GetAllSort{})
	s.Require().Error(err)
	s.Require().Empty(res)
}

func (s *UseCaseSuite) TestGetAllSuccess() {
	s.columnRepository.
		On("GetAll", s.ctx, mock.Anything, mock.Anything).
		Return([]domainColumn.Column{{Id: int(gofakeit.Int64())}}, nil).Once()

	res, err := s.useCase.GetAll(s.ctx, domainColumn.GetAllFilter{}, domainColumn.GetAllSort{})
	s.Require().NoError(err)
	s.Require().NotEmpty(res)
}
