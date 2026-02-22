package task

import (
	"context"

	taskConverter "github.com/poymanov/codemania-task-board/board/internal/infrastructure/converter/task"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetAll(ctx context.Context, req *boardV1.TaskServiceGetAllRequest) (*boardV1.TaskServiceGetAllResponse, error) {
	filter, sort := taskConverter.GetAllRequestToDomain(req)

	tasks, err := s.taskGetAllUseCase.GetAll(ctx, filter, sort)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to get all tasks")
		return nil, status.Error(codes.Internal, "failed to get all tasks")
	}

	responseTasks := make([]*boardV1.Task, 0, len(tasks))

	for _, task := range tasks {
		responseTasks = append(responseTasks, taskConverter.DomainToTransport(task))
	}

	return &boardV1.TaskServiceGetAllResponse{
		Tasks: responseTasks,
	}, nil
}
