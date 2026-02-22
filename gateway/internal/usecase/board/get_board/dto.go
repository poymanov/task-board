package get_board

import boardGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/board"

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

func ConvertGRPCClientDTOToDTO(board boardGrpcClientV1.BoardGetBoardDTO) BoardGetBoardDTO {
	columns := make([]ColumnGetBoardDTO, 0, len(board.Columns))

	for _, column := range board.Columns {
		tasks := make([]TaskGetBoardDTO, 0, len(column.Tasks))

		for _, task := range column.Tasks {
			tasks = append(tasks, TaskGetBoardDTO{
				Id:          task.Id,
				Title:       task.Title,
				Description: task.Description,
				Assignee:    task.Assignee,
				Position:    task.Position,
			})
		}

		columns = append(columns, ColumnGetBoardDTO{
			Id:       column.Id,
			Name:     column.Name,
			Position: column.Position,
			Tasks:    tasks,
		})
	}

	return BoardGetBoardDTO{
		Id:          board.Id,
		Name:        board.Name,
		Description: board.Description,
		OwnerId:     board.OwnerId,
		Columns:     columns,
	}
}
