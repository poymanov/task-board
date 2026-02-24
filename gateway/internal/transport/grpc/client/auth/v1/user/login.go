package user

import (
	"context"

	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
)

func (c *Client) Login(ctx context.Context, req LoginRequest) (string, error) {
	grpcReq := &authV1.UserServiceLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	resp, err := c.generatedClient.Login(ctx, grpcReq)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
