package common

import (
	domainCommon "github.com/poymanov/codemania-task-board/board/internal/domain/common"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func BoardDomainToTransport(board domainCommon.Board) *boardV1.BoardServiceGetBoardResponse {
	columns := make([]*boardV1.ColumnGetBoard, 0, len(board.Columns))

	for _, column := range board.Columns {
		tasks := make([]*boardV1.TaskGetBoard, 0, len(column.Tasks))

		for _, task := range column.Tasks {
			tasks = append(tasks, &boardV1.TaskGetBoard{
				Id:          int64(task.Id),
				Title:       task.Title,
				Description: task.Description,
				Assignee:    task.Assignee,
				Position:    float32(task.Position),
			})
		}

		columns = append(columns, &boardV1.ColumnGetBoard{
			Id:       int64(column.Id),
			Name:     column.Name,
			Position: float32(column.Position),
			Tasks:    tasks,
		})
	}

	return &boardV1.BoardServiceGetBoardResponse{
		Board: &boardV1.BoardGetBoard{
			Id:          int64(board.Id),
			Name:        board.Name,
			Description: board.Description,
			OwnerId:     int64(board.OwnerId),
			Columns:     columns,
		},
	}
}
