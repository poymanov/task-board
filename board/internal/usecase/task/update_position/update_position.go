package update_position

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

func (u *UseCase) UpdatePosition(ctx context.Context, id int, updatePositionDTO UpdatePositionDTO) error {
	exist, err := u.taskRepository.IsExistsById(ctx, id)
	if err != nil {
		return err
	}

	if !exist {
		return domainTask.ErrTaskNotExists
	}

	newPosition := (updatePositionDTO.LeftPosition + updatePositionDTO.RightPosition) / 2

	err = u.taskRepository.UpdatePosition(ctx, id, newPosition)
	if err != nil {
		return err
	}

	return nil
}
