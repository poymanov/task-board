package board

import (
	"context"

	"github.com/poymanov/codemania-task-board/platform/pkg/otel/tracer"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *BoardClient) GetBoard(ctx context.Context, request GetBoardRequest) (BoardGetBoardDTO, error) {
	ctx, span := tracer.Start(ctx, "GetBoard gRPC Request")
	defer span.End()

	grpcReq := &boardV1.BoardServiceGetBoardRequest{
		Id: int64(request.Id),
	}

	res, err := c.generatedClient.GetBoard(ctx, grpcReq)
	if err != nil {
		return BoardGetBoardDTO{}, err
	}

	return ConvertTransportGetBoardToDTO(res.GetBoard()), nil
}
