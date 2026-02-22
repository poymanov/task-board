package create

import (
	"context"

	"github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/board"
)

type UseCase struct {
	boardClient *board.BoardClient
}

func NewUseCase(boardClient *board.BoardClient) *UseCase {
	return &UseCase{
		boardClient: boardClient,
	}
}

func (u *UseCase) Create(ctx context.Context, dto CreateBoardDTO) (int, error) {
	createBoardRequest := board.CreateBoardRequest{Name: dto.Name, Description: dto.Description, OwnerId: dto.OwnerId}

	boardId, err := u.boardClient.CreateBoard(ctx, createBoardRequest)
	if err != nil {
		return 0, nil
	}

	return boardId, nil
}
