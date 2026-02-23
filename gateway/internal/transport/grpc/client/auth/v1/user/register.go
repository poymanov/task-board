package user

import (
	"context"

	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
)

func (c *Client) Register(ctx context.Context, req RegisterRequest) error {
	grpcReq := &authV1.UserServiceRegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}

	_, err := c.generatedClient.Register(ctx, grpcReq)
	if err != nil {
		return err
	}

	return nil
}
