package outbox_event

import (
	"github.com/jackc/pgx/v5"
	domainCommon "github.com/poymanov/codemania-task-board/board/internal/domain/common"
	domainOutboxEvent "github.com/poymanov/codemania-task-board/board/internal/domain/outbox_event"
)

type Repository struct {
	pool domainCommon.DB
}

func NewRepository(pool domainCommon.DB) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) WithTx(tx pgx.Tx) domainOutboxEvent.OutboxEventRepository {
	return &Repository{pool: tx}
}
