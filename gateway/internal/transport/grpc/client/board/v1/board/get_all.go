package board

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *BoardClient) GetAllBoard(ctx context.Context) ([]GetAllBoardDTO, error) {
	grpcReq := &boardV1.BoardServiceGetAllRequest{}

	res, err := c.generatedClient.GetAll(ctx, grpcReq)
	if err != nil {
		return []GetAllBoardDTO{}, err
	}

	responseBoards := res.GetBoards()

	dtos := make([]GetAllBoardDTO, 0, len(responseBoards))

	for _, board := range responseBoards {
		dtos = append(dtos, ConvertTransportGetAllBoardToDTO(board))
	}

	return dtos, nil
}
