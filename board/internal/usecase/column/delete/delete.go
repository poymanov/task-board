package delete

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

func (u *UseCase) Delete(ctx context.Context, id int) error {
	return u.columnRepository.Delete(ctx, id)
}
