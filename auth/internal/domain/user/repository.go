package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, newUser NewUser) error

	IsDuplicateKey(err error) bool
}
