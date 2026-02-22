package column

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
)

func (r *Repository) GetAll(ctx context.Context, filter domainColumn.GetAllFilter, sort domainColumn.GetAllSort) ([]domainColumn.Column, error) {
	query, args := getQuery(filter, sort)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return []domainColumn.Column{}, err
	}

	defer rows.Close()

	models, err := pgx.CollectRows(rows, pgx.RowToStructByName[Column])
	if err != nil {
		return []domainColumn.Column{}, err
	}

	columns := make([]domainColumn.Column, 0, len(models))

	for _, model := range models {
		columns = append(columns, ConventModelToDomain(model))
	}

	return columns, nil
}

func getQuery(filter domainColumn.GetAllFilter, sort domainColumn.GetAllSort) (string, []any) {
	var conditions []string
	var sorts []string
	var args []any
	argsPos := 1

	query := "SELECT * FROM columns"

	if filter.BoardId != 0 {
		conditions = append(conditions, fmt.Sprintf("board_id = $%d", argsPos))
		args = append(args, filter.BoardId)
	}

	if sort.SortByPosition != "" && (strings.ToLower(sort.SortByPosition) == "asc" || strings.ToLower(sort.SortByPosition) == "desc") {
		sorts = append(sorts, fmt.Sprintf("position %s", sort.SortByPosition))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if len(sorts) > 0 {
		query += " ORDER BY " + strings.Join(sorts, ", ")
	}

	return query, args
}
