package v1

import (
	boardGetAllUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/get_all"
	boardGetBoardUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/get_board"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
)

func GetAllBoardDTOToTransport(board boardGetAllUseCase.BoardDTO) gatewayV1.GetAllBoardResponseItem {
	return gatewayV1.GetAllBoardResponseItem{
		ID:          board.Id,
		Name:        board.Name,
		Description: board.Description,
	}
}

func GetBoardDTOToTransport(board boardGetBoardUseCase.BoardGetBoardDTO) *gatewayV1.GetBoardResponse {
	columns := make([]gatewayV1.GetBoardResponseColumnsItem, 0, len(board.Columns))

	for _, column := range board.Columns {
		tasks := make([]gatewayV1.GetBoardResponseColumnsItemTasksItem, 0, len(column.Tasks))

		for _, task := range column.Tasks {
			tasks = append(tasks, gatewayV1.GetBoardResponseColumnsItemTasksItem{
				ID:          task.Id,
				Title:       task.Title,
				Description: task.Description,
				Assignee:    task.Assignee,
				Position:    task.Position,
			})
		}

		columns = append(columns, gatewayV1.GetBoardResponseColumnsItem{
			ID:       column.Id,
			Name:     column.Name,
			Position: column.Position,
			Tasks:    tasks,
		})
	}

	return &gatewayV1.GetBoardResponse{
		ID:          board.Id,
		Name:        board.Name,
		Description: board.Description,
		Columns:     columns,
	}
}
