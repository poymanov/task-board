package task

import (
	"context"
	"errors"

	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	taskConverter "github.com/poymanov/codemania-task-board/board/internal/infrastructure/converter/task"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) UpdatePosition(ctx context.Context, req *boardV1.TaskServiceUpdatePositionRequest) (*boardV1.TaskServiceUpdatePositionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	id := int(req.GetId())

	dto := taskConverter.UpdatePositionRequestToUseCaseDTO(req)

	err := s.taskUpdatePositionUseCase.UpdatePosition(ctx, id, dto)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to update task position")

		if errors.Is(err, domainTask.ErrTaskNotExists) {
			return nil, status.Errorf(codes.Internal, "failed to update task position: %v", err)
		}

		return nil, status.Error(codes.Internal, "failed to update task position")
	}

	return &boardV1.TaskServiceUpdatePositionResponse{}, nil
}
