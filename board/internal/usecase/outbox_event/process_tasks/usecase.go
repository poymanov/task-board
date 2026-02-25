package process_tasks

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/jackc/pgx/v5"
	domainOutboxEvent "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event"
	"github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/tx_manager"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type UseCase struct {
	txManager tx_manager.Tx

	outboxEventRepository domainOutboxEvent.OutboxEventRepository

	taskChangedProducer *kafka.Writer
}

func NewUseCase(outboxEventRepository domainOutboxEvent.OutboxEventRepository, txManager tx_manager.Tx, taskChangedProducer *kafka.Writer) *UseCase {
	return &UseCase{
		outboxEventRepository: outboxEventRepository,
		txManager:             txManager,
		taskChangedProducer:   taskChangedProducer,
	}
}

func (u *UseCase) Process(ctx context.Context, limit int) error {
	tasks, err := u.outboxEventRepository.GetAllNotProcessedByType(ctx, "task", limit)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	wg.Add(len(tasks))
	for _, task := range tasks {
		dto := OutboxEventProducerDTO{
			Id:         task.Id,
			Payload:    task.Payload,
			EntityId:   task.EntityId,
			EntityType: task.EntityType,
			CreatedAt:  task.CreatedAt,
		}

		producerData, je := json.Marshal(dto)

		if je != nil {
			log.Error().Any("task", task).Err(je).Msg("failed to convert task to producer data")
			continue
		}

		go func() {
			wg.Done()
			etx := u.txManager.WithTx(ctx, func(tx pgx.Tx) error {
				outboxEventRepo := u.outboxEventRepository.WithTx(tx)

				eu := outboxEventRepo.UpdateEventProcessed(ctx, task.Id)

				if eu != nil {
					return eu
				}

				ew := u.taskChangedProducer.WriteMessages(ctx,
					kafka.Message{
						Value: producerData,
					},
				)
				if ew != nil {
					return ew
				}

				log.Info().Any("data", string(producerData)).Msg("Event processed")

				return nil
			})

			if etx != nil {
				log.Error().Any("task", task).Err(etx).Msg("failed to process event")
			}
		}()
	}

	wg.Wait()

	return nil
}
