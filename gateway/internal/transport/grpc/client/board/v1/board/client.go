package board

import boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"

type BoardClient struct {
	generatedClient boardV1.BoardServiceClient
}

func NewClient(client boardV1.BoardServiceClient) *BoardClient {
	return &BoardClient{client}
}
