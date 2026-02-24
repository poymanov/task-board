package user

import (
	"context"

	"github.com/jackc/pgx/v5"
	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
)

func (r *Repository) GetByEmail(ctx context.Context, email string) (domainUser.User, error) {
	row, err := r.pool.Query(ctx, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return domainUser.User{}, err
	}

	model, err := pgx.CollectOneRow(row, pgx.RowToStructByName[User])
	if err != nil {
		return domainUser.User{}, err
	}

	return ConvertModelToDomain(model), nil
}
