package column

import (
	"context"
	"time"
)

func (r *Repository) UpdatePosition(ctx context.Context, id int, position float64) error {
	_, err := r.pool.Exec(ctx, "UPDATE columns SET position=$1, updated_at=$2 WHERE id=$3", position, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
