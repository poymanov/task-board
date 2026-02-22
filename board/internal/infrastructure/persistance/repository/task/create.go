package task

import (
	"context"
	"fmt"

	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

func (r *Repository) Create(ctx context.Context, newTask domainTask.NewTask) (int, error) {
	var id int

	err := r.pool.QueryRow(
		ctx,
		`
		INSERT INTO tasks (title, description, assignee , position, column_id)
		VALUES (
   		$1,
		$2,
		$3,
   		(SELECT COALESCE(MAX(position), 0) + 1000
    	FROM tasks
    	WHERE column_id = $4),
   		$4
		) RETURNING id`,
		newTask.Title, newTask.Description, newTask.Assignee, newTask.ColumnId,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}

	return id, nil
}
