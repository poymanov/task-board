package register

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

func (u *UseCase) Register(ctx context.Context, dto RegisterDTO) error {
	registerRequest := user.RegisterRequest{Email: dto.Email, Password: dto.Password, Username: dto.Username}

	err := u.userClient.Register(ctx, registerRequest)
	if err != nil {
		return err
	}

	return nil
}
