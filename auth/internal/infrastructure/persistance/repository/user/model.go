package user

import (
	"time"

	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
)

type User struct {
	Id int `db:"id"`

	Username string `db:"username"`

	Email string `db:"email"`

	Password string `db:"password"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt *time.Time `db:"updated_at"`
}

func ConvertModelToDomain(user User) domainUser.User {
	return domainUser.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
