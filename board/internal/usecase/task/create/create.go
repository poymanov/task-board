package create

import (
	"context"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

type UseCase struct {
	columnRepository domainColumn.ColumnRepository

	taskRepository domainTask.TaskRepository
}

func NewUseCase(columnRepository domainColumn.ColumnRepository, taskRepository domainTask.TaskRepository) *UseCase {
	return &UseCase{
		columnRepository: columnRepository,
		taskRepository:   taskRepository,
	}
}

func (u *UseCase) Create(ctx context.Context, newTask NewTaskDTO) (int, error) {
	columnExists, err := u.columnRepository.IsExistsById(ctx, newTask.ColumnId)
	if err != nil {
		return 0, err
	}

	if !columnExists {
		return 0, domainColumn.ErrColumnNotExists
	}

	nt := domainTask.NewNewTask(newTask.Title, newTask.Description, newTask.Assignee, newTask.ColumnId)

	return u.taskRepository.Create(ctx, nt)
}
