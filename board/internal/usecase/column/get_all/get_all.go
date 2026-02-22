package get_all

import (
	"context"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
)

type UseCase struct {
	columnRepository domainColumn.ColumnRepository
}

func NewUseCase(columnRepository domainColumn.ColumnRepository) *UseCase {
	return &UseCase{
		columnRepository: columnRepository,
	}
}

func (u *UseCase) GetAll(ctx context.Context, filter domainColumn.GetAllFilter, sort domainColumn.GetAllSort) ([]domainColumn.Column, error) {
	return u.columnRepository.GetAll(ctx, filter, sort)
}
