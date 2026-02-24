package register

import (
	"context"

	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	userRepository domainUser.UserRepository
}

func NewUseCase(userRepository domainUser.UserRepository) *UseCase {
	return &UseCase{
		userRepository: userRepository,
	}
}

func (u *UseCase) Register(ctx context.Context, registerUser RegisterUserDTO) error {
	password, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	nu := domainUser.NewNewUser(registerUser.Email, string(password), registerUser.Username)

	err = u.userRepository.Create(ctx, nu)
	if err != nil {
		if u.userRepository.IsDuplicateKey(err) {
			return domainUser.ErrUserAlreadyExists
		}

		return err
	}

	return nil
}
