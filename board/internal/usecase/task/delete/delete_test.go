package delete

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestDeleteError() {
	id := int(gofakeit.Int64())

	errDelete := errors.New(gofakeit.Word())

	s.txManager.
		On("WithTx", mock.Anything, mock.Anything).
		Return(errDelete).
		Once()

	err := s.useCase.Delete(s.ctx, id)
	s.Require().Error(err)
}

func (s *UseCaseSuite) TestDeleteSuccess() {
	id := int(gofakeit.Int64())

	s.txManager.
		On("WithTx", mock.Anything, mock.Anything).
		Return(nil).
		Once()

	err := s.useCase.Delete(s.ctx, id)
	s.Require().NoError(err)
}
