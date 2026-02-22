package board

import (
	boardCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/create"
	boardDeleteUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/delete"
	boardGetAllUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/get_all"
	boardGetBoardUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board/get_board"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

type BoardService struct {
	boardCreateUseCase   *boardCreateUseCase.UseCase
	boardGetAllUseCase   *boardGetAllUseCase.UseCase
	boardDeleteUseCase   *boardDeleteUseCase.UseCase
	boardGetBoardUseCase *boardGetBoardUseCase.UseCase

	boardV1.UnimplementedBoardServiceServer
}

func NewBoardService(
	boardCreateUseCase *boardCreateUseCase.UseCase,
	boardGetAllUseCase *boardGetAllUseCase.UseCase,
	boardDeleteUseCase *boardDeleteUseCase.UseCase,
	boardGetBoardUseCase *boardGetBoardUseCase.UseCase,
) *BoardService {
	return &BoardService{
		boardCreateUseCase:   boardCreateUseCase,
		boardGetAllUseCase:   boardGetAllUseCase,
		boardDeleteUseCase:   boardDeleteUseCase,
		boardGetBoardUseCase: boardGetBoardUseCase,
	}
}
