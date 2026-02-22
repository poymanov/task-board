package task

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

func (r *Repository) GetAll(ctx context.Context, filter domainTask.GetAllFilter, sort domainTask.GetAllSort) ([]domainTask.Task, error) {
	query, args := getQuery(filter, sort)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return []domainTask.Task{}, err
	}

	defer rows.Close()

	models, err := pgx.CollectRows(rows, pgx.RowToStructByName[Task])
	if err != nil {
		return []domainTask.Task{}, err
	}

	tasks := make([]domainTask.Task, 0, len(models))

	for _, model := range models {
		tasks = append(tasks, ConventModelToDomain(model))
	}

	return tasks, nil
}

func getQuery(filter domainTask.GetAllFilter, sort domainTask.GetAllSort) (string, []any) {
	var conditions []string
	var sorts []string
	var args []any
	argsPos := 1

	query := "SELECT * FROM tasks"

	if len(filter.ColumnIds) > 0 {
		conditions = append(conditions, fmt.Sprintf("column_id = ANY($%d)", argsPos))
		args = append(args, filter.ColumnIds)
	}

	if sort.SortByColumnId != "" && (strings.ToLower(sort.SortByColumnId) == "asc" || strings.ToLower(sort.SortByColumnId) == "desc") {
		sorts = append(sorts, fmt.Sprintf("column_id %s", sort.SortByColumnId))
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
