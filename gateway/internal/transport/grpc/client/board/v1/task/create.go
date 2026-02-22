package task

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *Client) Create(ctx context.Context, req CreateTaskRequest) (int, error) {
	grpcReq := &boardV1.TaskServiceCreateRequest{
		Title:       req.Title,
		Description: req.Description,
		Assignee:    req.Assignee,
		ColumnId:    int64(req.ColumnId),
	}

	res, err := c.generatedClient.Create(ctx, grpcReq)
	if err != nil {
		return 0, err
	}

	return int(res.GetTaskId()), nil
}
