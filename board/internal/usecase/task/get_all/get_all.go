package get_all

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

func (u *UseCase) GetAll(ctx context.Context, filter domainTask.GetAllFilter, sort domainTask.GetAllSort) ([]domainTask.Task, error) {
	return u.taskRepository.GetAll(ctx, filter, sort)
}
