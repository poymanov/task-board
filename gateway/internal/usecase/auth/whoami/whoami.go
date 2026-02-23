package whoami

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

func (u *UseCase) Whoami(ctx context.Context, dto WhoamiDTO) (WhoamiUserDTO, error) {
	whoamiRequest := user.WhoamiRequest{AccessToken: dto.AccessToken}

	whoami, err := u.userClient.Whoami(ctx, whoamiRequest)
	if err != nil {
		return WhoamiUserDTO{}, err
	}

	userDto := WhoamiUserDTO{
		Username: whoami.Username,
		Email:    whoami.Email,
		UserId:   whoami.UserId,
	}

	return userDto, nil
}
