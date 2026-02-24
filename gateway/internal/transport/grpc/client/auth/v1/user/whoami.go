package user

import (
	"context"

	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
)

func (c *Client) Whoami(ctx context.Context, req WhoamiRequest) (WhoamiDTO, error) {
	grpcReq := &authV1.UserServiceWhoamiRequest{
		AccessToken: req.AccessToken,
	}

	resp, err := c.generatedClient.Whoami(ctx, grpcReq)
	if err != nil {
		return WhoamiDTO{}, err
	}

	dto := WhoamiDTO{
		UserId:   int(resp.GetUserId()),
		Email:    resp.GetEmail(),
		Username: resp.GetUsername(),
	}

	return dto, nil
}
