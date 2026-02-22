package column

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *Client) Create(ctx context.Context, req CreateColumnRequest) (int, error) {
	grpcReq := &boardV1.ColumnServiceCreateRequest{
		Name:    req.Name,
		BoardId: int64(req.BoardId),
	}

	res, err := c.generatedClient.Create(ctx, grpcReq)
	if err != nil {
		return 0, err
	}

	return int(res.GetColumnId()), nil
}
