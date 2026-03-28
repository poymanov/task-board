package task

import (
	"context"
)

func (r *Repository) Delete(ctx context.Context, id int) (bool, error) {
	res, err := r.pool.Exec(ctx, "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return false, err
	}

	if res.RowsAffected() > 0 {
		return true, nil
	}

	return false, nil
}
