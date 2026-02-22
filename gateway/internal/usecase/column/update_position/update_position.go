package update_position

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

func (u *UseCase) UpdatePosition(ctx context.Context, dto UpdatePositionColumnDTO) error {
	updatePositionColumnRequest := column.UpdatePositionColumnRequest{Id: dto.Id, LeftPosition: dto.LeftPosition, RightPosition: dto.RightPosition}

	err := u.columnClient.UpdatePosition(ctx, updatePositionColumnRequest)
	if err != nil {
		return err
	}

	return nil
}
