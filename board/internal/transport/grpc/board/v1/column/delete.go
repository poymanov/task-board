package column

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Delete(ctx context.Context, req *boardV1.ColumnServiceDeleteRequest) (*boardV1.ColumnServiceDeleteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	id := int(req.GetId())

	err := s.columnDeleteUseCase.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Any("id", id).Msg("failed to delete column")
		return nil, status.Error(codes.Internal, "failed to delete column")
	}

	return &boardV1.ColumnServiceDeleteResponse{}, nil
}
