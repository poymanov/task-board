package column

import boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"

type Client struct {
	generatedClient boardV1.ColumnServiceClient
}

func NewClient(client boardV1.ColumnServiceClient) *Client {
	return &Client{client}
}
