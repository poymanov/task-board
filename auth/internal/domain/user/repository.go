package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, newUser NewUser) error

	GetByEmail(ctx context.Context, email string) (User, error)

	IsDuplicateKey(err error) bool

	IsNoRows(err error) bool
}
