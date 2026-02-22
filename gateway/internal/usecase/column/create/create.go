package create

import (
	"context"

	"github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/column"
)

type UseCase struct {
	columnClient *column.Client
}

func NewUseCase(columnClient *column.Client) *UseCase {
	return &UseCase{
		columnClient: columnClient,
	}
}

func (u *UseCase) Create(ctx context.Context, dto CreateColumnDTO) (int, error) {
	createColumnRequest := column.CreateColumnRequest{Name: dto.Name, BoardId: dto.BoardId}

	boardId, err := u.columnClient.Create(ctx, createColumnRequest)
	if err != nil {
		return 0, err
	}

	return boardId, nil
}
