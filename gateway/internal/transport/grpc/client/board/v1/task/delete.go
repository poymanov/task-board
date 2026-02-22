package task

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *Client) Delete(ctx context.Context, req DeleteTaskRequest) error {
	grpcReq := &boardV1.TaskServiceDeleteRequest{
		Id: int64(req.Id),
	}

	_, err := c.generatedClient.Delete(ctx, grpcReq)
	if err != nil {
		return err
	}

	return nil
}
