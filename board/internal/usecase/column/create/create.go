package create

import (
	"context"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
)

type UseCase struct {
	boardRepository  domainBoard.BoardRepository
	columnRepository domainColumn.ColumnRepository
}

func NewUseCase(boardRepository domainBoard.BoardRepository, columnRepository domainColumn.ColumnRepository) *UseCase {
	return &UseCase{
		boardRepository:  boardRepository,
		columnRepository: columnRepository,
	}
}

func (u *UseCase) Create(ctx context.Context, newColumn NewColumnDTO) (int, error) {
	boardExists, err := u.boardRepository.IsExistsById(ctx, newColumn.BoardID)
	if err != nil {
		return 0, err
	}

	if !boardExists {
		return 0, domainBoard.ErrBoardNotExists
	}

	nc := domainColumn.NewNewColumn(newColumn.Name, newColumn.BoardID)

	return u.columnRepository.Create(ctx, nc)
}
