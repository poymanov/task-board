package board

import boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"

type GetAllBoardDTO struct {
	Id int

	Name string

	Description string
}

type BoardGetBoardDTO struct {
	Id int

	Name string

	Description string

	OwnerId int

	Columns []ColumnGetBoardDTO
}

type ColumnGetBoardDTO struct {
	Id int

	Name string

	Position float64

	Tasks []TaskGetBoardDTO
}

type TaskGetBoardDTO struct {
	Id int

	Title string

	Description string

	Assignee string

	Position float64
}

func ConvertTransportGetAllBoardToDTO(board *boardV1.Board) GetAllBoardDTO {
	return GetAllBoardDTO{
		Id:          int(board.Id),
		Name:        board.Name,
		Description: board.Description,
	}
}

func ConvertTransportGetBoardToDTO(board *boardV1.BoardGetBoard) BoardGetBoardDTO {
	columns := make([]ColumnGetBoardDTO, 0, len(board.Columns))

	for _, column := range board.Columns {
		tasks := make([]TaskGetBoardDTO, 0, len(column.Tasks))

		for _, task := range column.Tasks {
			tasks = append(tasks, TaskGetBoardDTO{
				Id:          int(task.Id),
				Title:       task.Title,
				Description: task.Description,
				Assignee:    task.Assignee,
				Position:    float64(task.Position),
			})
		}

		columns = append(columns, ColumnGetBoardDTO{
			Id:       int(column.Id),
			Name:     column.Name,
			Position: float64(column.Position),
			Tasks:    tasks,
		})
	}

	return BoardGetBoardDTO{
		Id:          int(board.Id),
		Name:        board.Name,
		Description: board.Description,
		OwnerId:     int(board.OwnerId),
		Columns:     columns,
	}
}
