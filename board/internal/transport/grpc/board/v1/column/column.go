package column

import (
	columnCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/create"
	columnDeleteUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/delete"
	columnGetAllUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/get_all"
	columnUpdatePositionUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/update_position"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

type Service struct {
	columnCreateUseCase         *columnCreateUseCase.UseCase
	columnGetAllUseCase         *columnGetAllUseCase.UseCase
	columnDeleteUseCase         *columnDeleteUseCase.UseCase
	columnUpdatePositionUseCase *columnUpdatePositionUseCase.UseCase

	boardV1.UnimplementedColumnServiceServer
}

func NewService(
	columnCreateUseCase *columnCreateUseCase.UseCase,
	columnGetAllUseCase *columnGetAllUseCase.UseCase,
	columnDeleteUseCase *columnDeleteUseCase.UseCase,
	columnUpdatePositionUseCase *columnUpdatePositionUseCase.UseCase,
) *Service {
	return &Service{
		columnCreateUseCase:         columnCreateUseCase,
		columnGetAllUseCase:         columnGetAllUseCase,
		columnDeleteUseCase:         columnDeleteUseCase,
		columnUpdatePositionUseCase: columnUpdatePositionUseCase,
	}
}
