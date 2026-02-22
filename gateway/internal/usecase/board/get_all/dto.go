package get_all

import (
	boardGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/board"
)

type BoardDTO struct {
	Id int

	Name string

	Description string
}

func ConvertGRPCClientDTOToDTO(board boardGrpcClientV1.GetAllBoardDTO) BoardDTO {
	return BoardDTO{
		Id:          board.Id,
		Name:        board.Name,
		Description: board.Description,
	}
}
