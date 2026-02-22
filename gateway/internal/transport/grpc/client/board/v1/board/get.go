package board

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *BoardClient) GetBoard(ctx context.Context, request GetBoardRequest) (BoardGetBoardDTO, error) {
	grpcReq := &boardV1.BoardServiceGetBoardRequest{
		Id: int64(request.Id),
	}

	res, err := c.generatedClient.GetBoard(ctx, grpcReq)
	if err != nil {
		return BoardGetBoardDTO{}, err
	}

	return ConvertTransportGetBoardToDTO(res.GetBoard()), nil
}
