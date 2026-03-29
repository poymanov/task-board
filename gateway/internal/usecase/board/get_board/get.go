package get_board

import (
	"context"

	boardGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/board"
	"github.com/poymanov/codemania-task-board/platform/pkg/otel/tracer"
)

type UseCase struct {
	boardClient *boardGrpcClientV1.BoardClient
}

func NewUseCase(boardClient *boardGrpcClientV1.BoardClient) *UseCase {
	return &UseCase{
		boardClient: boardClient,
	}
}

func (u *UseCase) Get(ctx context.Context, id int) (BoardGetBoardDTO, error) {
	ctx, span := tracer.Start(ctx, "GetBoard useCase")
	defer span.End()

	getBoardRequest := boardGrpcClientV1.GetBoardRequest{Id: id}

	getBoardDto, err := u.boardClient.GetBoard(ctx, getBoardRequest)
	if err != nil {
		return BoardGetBoardDTO{}, err
	}

	return ConvertGRPCClientDTOToDTO(getBoardDto), nil
}
