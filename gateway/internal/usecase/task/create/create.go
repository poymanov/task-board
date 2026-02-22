package create

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

func (u *UseCase) Create(ctx context.Context, dto TaskCreateDTO) (int, error) {
	createTaskRequest := task.CreateTaskRequest{Title: dto.Title, Description: dto.Description, Assignee: dto.Assignee, ColumnId: dto.ColumnId}

	boardId, err := u.taskClient.Create(ctx, createTaskRequest)
	if err != nil {
		return 0, err
	}

	return boardId, nil
}
