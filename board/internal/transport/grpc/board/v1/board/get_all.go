package board

import (
	"context"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	boardConverter "github.com/poymanov/codemania-task-board/board/internal/infrastructure/converter/board"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BoardService) GetAll(ctx context.Context, req *boardV1.BoardServiceGetAllRequest) (*boardV1.BoardServiceGetAllResponse, error) {
	ownerId := int(req.GetFilter().GetOwnerId())

	filter := domainBoard.NewGetAllFilter(ownerId)

	boards, err := s.boardGetAllUseCase.GetAll(ctx, filter)
	if err != nil {
		log.Error().Err(err).Any("filter", filter).Msg("failed to get all boards")
		return nil, status.Errorf(codes.Internal, "error getting all boards: %v", err)
	}

	responseBoards := make([]*boardV1.Board, 0, len(boards))

	for _, board := range boards {
		responseBoards = append(responseBoards, boardConverter.DomainToTransport(board))
	}

	return &boardV1.BoardServiceGetAllResponse{
		Boards: responseBoards,
	}, nil
}
