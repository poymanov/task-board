package task

import (
	taskCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/create"
	taskDeleteUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/delete"
	taskGetAllUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/get_all"
	taskUpdatePositionUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/update_position"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

type Service struct {
	taskCreateUseCase *taskCreateUseCase.UseCase

	taskGetAllUseCase *taskGetAllUseCase.UseCase

	taskDeleteUseCase *taskDeleteUseCase.UseCase

	taskUpdatePositionUseCase *taskUpdatePositionUseCase.UseCase

	boardV1.UnimplementedTaskServiceServer
}

func NewService(
	taskCreateUseCase *taskCreateUseCase.UseCase,
	taskGetAllUseCase *taskGetAllUseCase.UseCase,
	taskDeleteUseCase *taskDeleteUseCase.UseCase,
	taskUpdatePositionUseCase *taskUpdatePositionUseCase.UseCase,
) *Service {
	return &Service{
		taskCreateUseCase:         taskCreateUseCase,
		taskGetAllUseCase:         taskGetAllUseCase,
		taskDeleteUseCase:         taskDeleteUseCase,
		taskUpdatePositionUseCase: taskUpdatePositionUseCase,
	}
}
