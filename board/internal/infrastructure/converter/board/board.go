package board

import (
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func DomainToTransport(board domainBoard.Board) *boardV1.Board {
	return &boardV1.Board{
		Id:          int64(board.Id),
		Name:        board.Name,
		Description: board.Description,
		OwnerId:     int64(board.OwnerId),
	}
}
