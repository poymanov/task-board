package delete

import (
	"context"

	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

type UseCase struct {
	taskRepository domainTask.TaskRepository
}

func NewUseCase(taskRepository domainTask.TaskRepository) *UseCase {
	return &UseCase{
		taskRepository: taskRepository,
	}
}

func (u *UseCase) Delete(ctx context.Context, id int) error {
	return u.taskRepository.Delete(ctx, id)
}
