package board

import (
	"context"

	"github.com/jackc/pgx/v5"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	"github.com/poymanov/codemania-task-board/platform/pkg/otel/tracer"
)

func (r *Repository) GetById(ctx context.Context, id int) (domainBoard.Board, error) {
	ctx, span := tracer.Start(ctx, "GetBoard repository GetById")
	defer span.End()

	row, err := r.pool.Query(ctx, "SELECT * FROM boards WHERE id=$1", id)
	if err != nil {
		return domainBoard.Board{}, domainBoard.ErrBoardNotExists
	}

	model, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Board])
	if err != nil {
		return domainBoard.Board{}, domainBoard.ErrBoardNotExists
	}

	return ConvertModelToDomain(model), nil
}
