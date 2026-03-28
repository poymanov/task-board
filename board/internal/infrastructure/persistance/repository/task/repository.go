package task

import (
	"github.com/jackc/pgx/v5"
	domainCommon "github.com/poymanov/codemania-task-board/board/internal/domain/common"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

type Repository struct {
	pool domainCommon.DB
}

func NewRepository(pool domainCommon.DB) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) WithTx(tx pgx.Tx) domainTask.TaskRepository {
	return &Repository{pool: tx}
}
