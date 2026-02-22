package board

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *BoardClient) CreateBoard(ctx context.Context, req CreateBoardRequest) (int, error) {
	grpcReq := &boardV1.BoardServiceCreateRequest{
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     int64(req.OwnerId),
	}

	res, err := c.generatedClient.Create(ctx, grpcReq)
	if err != nil {
		return 0, err
	}

	return int(res.GetBoardId()), nil
}
