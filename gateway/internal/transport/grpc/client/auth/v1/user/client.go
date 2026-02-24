package user

import (
	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
)

type Client struct {
	generatedClient authV1.UserServiceClient
}

func NewClient(client authV1.UserServiceClient) *Client {
	return &Client{client}
}
