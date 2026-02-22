package delete

import (
	"context"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

type UseCase struct {
	boardRepository domainBoard.BoardRepository
}

func NewUseCase(boardRepository domainBoard.BoardRepository) *UseCase {
	return &UseCase{
		boardRepository: boardRepository,
	}
}

func (u *UseCase) Delete(ctx context.Context, id int) error {
	return u.boardRepository.Delete(ctx, id)
}
