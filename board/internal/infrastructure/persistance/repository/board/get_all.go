package board

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

func (r *Repository) GetAll(ctx context.Context, filter domainBoard.GetAllFilter) ([]domainBoard.Board, error) {
	query, args := getQuery(filter)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return []domainBoard.Board{}, err
	}

	defer rows.Close()

	boardModels, err := pgx.CollectRows(rows, pgx.RowToStructByName[Board])
	if err != nil {
		return []domainBoard.Board{}, err
	}

	boards := make([]domainBoard.Board, 0, len(boardModels))

	for _, model := range boardModels {
		boards = append(boards, ConvertModelToDomain(model))
	}

	return boards, nil
}

func getQuery(filter domainBoard.GetAllFilter) (string, []any) {
	var conditions []string
	var args []any
	argsPos := 1

	query := "SELECT * FROM boards"

	if filter.OwnerId != 0 {
		conditions = append(conditions, fmt.Sprintf("owner_id = $%d", argsPos))
		args = append(args, filter.OwnerId)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return query, args
}
