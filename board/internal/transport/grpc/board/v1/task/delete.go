package task

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Delete(ctx context.Context, req *boardV1.TaskServiceDeleteRequest) (*boardV1.TaskServiceDeleteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	id := int(req.GetId())

	err := s.taskDeleteUseCase.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to delete task")
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &boardV1.TaskServiceDeleteResponse{}, nil
}
