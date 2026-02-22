package task

import (
	"context"
	"errors"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	taskCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/create"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Create(ctx context.Context, req *boardV1.TaskServiceCreateRequest) (*boardV1.TaskServiceCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	newTask := taskCreateUseCase.NewTaskDTO{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Assignee:    req.GetAssignee(),
		ColumnId:    int(req.GetColumnId()),
	}

	id, err := s.taskCreateUseCase.Create(ctx, newTask)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to create task")

		if errors.Is(err, domainColumn.ErrColumnNotExists) {
			return nil, status.Errorf(codes.Internal, "failed to create task: %v", err)
		}

		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &boardV1.TaskServiceCreateResponse{TaskId: int64(id)}, nil
}
