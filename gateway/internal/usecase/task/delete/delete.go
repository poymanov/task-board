package delete

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

func (u *UseCase) Delete(ctx context.Context, id int) error {
	taskDeleteRequest := task.DeleteTaskRequest{Id: id}

	err := u.taskClient.Delete(ctx, taskDeleteRequest)
	if err != nil {
		return err
	}

	return nil
}
