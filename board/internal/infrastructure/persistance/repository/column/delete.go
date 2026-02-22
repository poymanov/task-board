package column

import (
	"context"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM columns WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
