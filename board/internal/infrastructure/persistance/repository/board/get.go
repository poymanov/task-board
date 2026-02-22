package board

import (
	"context"

	"github.com/jackc/pgx/v5"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

func (r *Repository) GetById(ctx context.Context, id int) (domainBoard.Board, error) {
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
