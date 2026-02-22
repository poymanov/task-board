package board

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BoardService) Delete(ctx context.Context, req *boardV1.BoardServiceDeleteRequest) (*boardV1.BoardServiceDeleteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	id := int(req.GetId())

	err := s.boardDeleteUseCase.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Any("id", id).Msg("failed to delete board")
		return nil, status.Errorf(codes.Internal, "error delete board: %v", err)
	}

	return &boardV1.BoardServiceDeleteResponse{}, nil
}
