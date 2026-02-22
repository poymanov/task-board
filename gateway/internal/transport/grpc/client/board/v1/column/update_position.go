package column

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *Client) UpdatePosition(ctx context.Context, req UpdatePositionColumnRequest) error {
	grpcReq := &boardV1.ColumnServiceUpdatePositionRequest{
		Id:            int64(req.Id),
		LeftPosition:  float32(req.LeftPosition),
		RightPosition: float32(req.RightPosition),
	}

	_, err := c.generatedClient.UpdatePosition(ctx, grpcReq)
	if err != nil {
		return err
	}

	return nil
}
