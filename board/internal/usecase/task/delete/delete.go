package delete

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

func (u *UseCase) Delete(ctx context.Context, id int) error {
	return u.txManager.WithTx(ctx, func(tx pgx.Tx) error {
		taskRepo := u.taskRepository.WithTx(tx)

		deleted, err := taskRepo.Delete(ctx, id)
		if err != nil {
			return err
		}

		if !deleted {
			return nil
		}

		outboxEventRepo := u.outboxEventRepository.WithTx(tx)

		eventPayload := map[string]string{
			"event": "delete",
		}

		event := domainOutboxEvent.NewNewEvent("task", id, eventPayload)

		err = outboxEventRepo.Create(ctx, event)
		if err != nil {
			return err
		}

		return nil
	})
}
