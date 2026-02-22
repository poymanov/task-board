package column

import (
	"context"
	"fmt"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
)

func (r *Repository) Create(ctx context.Context, newColumn domainColumn.NewColumn) (int, error) {
	var id int

	err := r.pool.QueryRow(
		ctx,
		`
		INSERT INTO columns (name, position, board_id)
		VALUES (
   		$1,
   		(SELECT COALESCE(MAX(position), 0) + 1000
    	FROM columns
    	WHERE board_id = $2),
   		$2
		) RETURNING id`,
		newColumn.Name, newColumn.BoardID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create column: %w", err)
	}

	return id, nil
}
