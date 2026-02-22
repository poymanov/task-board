package update_position

import (
	"context"

	"github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/task"
)

type UseCase struct {
	taskClient *task.Client
}

func NewUseCase(taskClient *task.Client) *UseCase {
	return &UseCase{
		taskClient: taskClient,
	}
}

func (u *UseCase) UpdatePosition(ctx context.Context, dto UpdatePositionColumnDTO) error {
	updatePositionTaskRequest := task.UpdatePositionTaskRequest{Id: dto.Id, LeftPosition: dto.LeftPosition, RightPosition: dto.RightPosition}

	err := u.taskClient.UpdatePosition(ctx, updatePositionTaskRequest)
	if err != nil {
		return err
	}

	return nil
}
