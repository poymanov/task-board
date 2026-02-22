package task

import (
	"context"
)

func (r *Repository) IsExistsById(ctx context.Context, id int) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM tasks WHERE id=$1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
