package login

import (
	"context"

	"github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/auth/v1/user"
)

type UseCase struct {
	userClient *user.Client
}

func NewUseCase(userClient *user.Client) *UseCase {
	return &UseCase{
		userClient: userClient,
	}
}

func (u *UseCase) Login(ctx context.Context, dto LoginDTO) (string, error) {
	loginRequest := user.LoginRequest{Email: dto.Email, Password: dto.Password}

	accessToken, err := u.userClient.Login(ctx, loginRequest)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
