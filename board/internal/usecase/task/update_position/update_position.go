package update_position

import (
	"context"

	"github.com/jackc/pgx/v5"
	domainOutboxEvent "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	"github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/tx_manager"
)

type UseCase struct {
	txManager tx_manager.Tx

	taskRepository domainTask.TaskRepository

	outboxEventRepository domainOutboxEvent.OutboxEventRepository
}

func NewUseCase(
	taskRepository domainTask.TaskRepository,
	outboxEventRepository domainOutboxEvent.OutboxEventRepository,
	txManager tx_manager.Tx,
) *UseCase {
	return &UseCase{
		taskRepository:        taskRepository,
		outboxEventRepository: outboxEventRepository,
		txManager:             txManager,
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

	return u.txManager.WithTx(ctx, func(tx pgx.Tx) error {
		taskRepo := u.taskRepository.WithTx(tx)

		newPosition := (updatePositionDTO.LeftPosition + updatePositionDTO.RightPosition) / 2

		err = taskRepo.UpdatePosition(ctx, id, newPosition)
		if err != nil {
			return err
		}

		outboxEventRepo := u.outboxEventRepository.WithTx(tx)

		eventPayload := map[string]string{
			"event": "update_position",
		}

		event := domainOutboxEvent.NewNewEvent("task", id, eventPayload)

		err = outboxEventRepo.Create(ctx, event)
		if err != nil {
			return err
		}

		return nil
	})
}
