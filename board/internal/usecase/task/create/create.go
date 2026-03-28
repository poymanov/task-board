package create

import (
	"context"

	"github.com/jackc/pgx/v5"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	domainOutboxEvent "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	"github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/tx_manager"
)

type UseCase struct {
	txManager tx_manager.Tx

	columnRepository domainColumn.ColumnRepository

	outboxEventRepository domainOutboxEvent.OutboxEventRepository

	taskRepository domainTask.TaskRepository
}

func NewUseCase(columnRepository domainColumn.ColumnRepository, taskRepository domainTask.TaskRepository, outboxEventRepository domainOutboxEvent.OutboxEventRepository, txManager tx_manager.Tx) *UseCase {
	return &UseCase{
		columnRepository:      columnRepository,
		taskRepository:        taskRepository,
		outboxEventRepository: outboxEventRepository,
		txManager:             txManager,
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

	var id int

	err = u.txManager.WithTx(ctx, func(tx pgx.Tx) error {
		taskRepo := u.taskRepository.WithTx(tx)

		id, err = taskRepo.Create(ctx, nt)
		if err != nil {
			return err
		}

		outboxEventRepo := u.outboxEventRepository.WithTx(tx)

		eventPayload := map[string]string{
			"event": "create",
		}

		event := domainOutboxEvent.NewNewEvent("task", id, eventPayload)

		err = outboxEventRepo.Create(ctx, event)
		if err != nil {
			return err
		}

		return nil
	})

	return id, err
}
