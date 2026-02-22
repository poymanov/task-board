package task

import boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"

type Client struct {
	generatedClient boardV1.TaskServiceClient
}

func NewClient(client boardV1.TaskServiceClient) *Client {
	return &Client{client}
}
