package get_board

import (
	"context"

	boardGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/board"
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
	getBoardRequest := boardGrpcClientV1.GetBoardRequest{Id: id}

	getBoardDto, err := u.boardClient.GetBoard(ctx, getBoardRequest)
	if err != nil {
		return BoardGetBoardDTO{}, err
	}

	return ConvertGRPCClientDTOToDTO(getBoardDto), nil
}
