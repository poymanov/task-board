package get_all

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

func (u *UseCase) GetAll(ctx context.Context, filter domainBoard.GetAllFilter) ([]domainBoard.Board, error) {
	return u.boardRepository.GetAll(ctx, filter)
}
