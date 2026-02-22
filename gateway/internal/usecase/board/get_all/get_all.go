package get_all

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

func (u *UseCase) GetAll(ctx context.Context) ([]BoardDTO, error) {
	getAllBoardDtos, err := u.boardClient.GetAllBoard(ctx)
	if err != nil {
		return []BoardDTO{}, nil
	}

	dtos := make([]BoardDTO, 0, len(getAllBoardDtos))

	for _, getAllBoardDto := range getAllBoardDtos {
		dtos = append(dtos, ConvertGRPCClientDTOToDTO(getAllBoardDto))
	}

	return dtos, nil
}
