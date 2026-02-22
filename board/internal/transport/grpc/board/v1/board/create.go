package board

import (
	"context"

	boardUsecase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/create"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BoardService) Create(ctx context.Context, req *boardV1.BoardServiceCreateRequest) (*boardV1.BoardServiceCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	newBoard := boardUsecase.NewBoardDTO{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		OwnerID:     int(req.GetOwnerId()),
	}

	boardId, err := s.boardCreateUseCase.Create(ctx, newBoard)
	if err != nil {
		log.Error().Err(err).Msg("failed to create board")
		return nil, status.Errorf(codes.Internal, "error creating board: %v", err)
	}

	return &boardV1.BoardServiceCreateResponse{BoardId: int64(boardId)}, nil
}
