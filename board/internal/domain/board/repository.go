package board

import "context"

type BoardRepository interface {
	Create(ctx context.Context, newBoard NewBoard) (int, error)

	GetAll(ctx context.Context, filter GetAllFilter) ([]Board, error)

	Delete(ctx context.Context, id int) error

	IsExistsById(ctx context.Context, ID int) (bool, error)

	GetById(ctx context.Context, id int) (Board, error)
}
