package login

import (
	"context"
	"errors"

	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	"github.com/poymanov/codemania-task-board/auth/internal/infrastructure/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	userRepository domainUser.UserRepository
	jwtService     jwt.JWT
}

func NewUseCase(userRepository domainUser.UserRepository, jwtService jwt.JWT) *UseCase {
	return &UseCase{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (u *UseCase) Login(ctx context.Context, login LoginDTO) (string, error) {
	user, err := u.userRepository.GetByEmail(ctx, login.Email)
	if err != nil {
		if u.userRepository.IsNoRows(err) {
			return "", errors.Join(err, domainUser.ErrInvalidCredentials)
		}

		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return "", domainUser.ErrInvalidCredentials
	}

	accessToken, err := u.jwtService.GenerateAccessToken(user)
	if err != nil {
		return "", errors.Join(err, domainUser.ErrInvalidCredentials)
	}

	return accessToken, nil
}
