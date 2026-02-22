package column

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *Client) Delete(ctx context.Context, req DeleteColumnRequest) error {
	grpcReq := &boardV1.ColumnServiceDeleteRequest{
		Id: int64(req.Id),
	}

	_, err := c.generatedClient.Delete(ctx, grpcReq)
	if err != nil {
		return err
	}

	return nil
}
