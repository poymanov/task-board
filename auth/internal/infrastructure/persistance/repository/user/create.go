package user

import (
	"context"

	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
)

func (r *Repository) Create(ctx context.Context, newUser domainUser.NewUser) error {
	_, err := r.pool.Exec(
		ctx,
		`INSERT INTO users (username, email, password) VALUES ($1,$2,$3)`,
		newUser.Username, newUser.Email, newUser.Password,
	)
	if err != nil {
		return err
	}

	return nil
}
