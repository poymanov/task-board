package column

import "context"

type ColumnRepository interface {
	Create(ctx context.Context, newColumn NewColumn) (int, error)

	GetAll(ctx context.Context, filter GetAllFilter, sort GetAllSort) ([]Column, error)

	Delete(ctx context.Context, id int) error

	UpdatePosition(ctx context.Context, id int, position float64) error

	IsExistsById(ctx context.Context, id int) (bool, error)
}
