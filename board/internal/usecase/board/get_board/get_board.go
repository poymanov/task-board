package get_board

import (
	"context"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	domainCommon "github.com/poymanov/codemania-task-board/board/internal/domain/common"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

type UseCase struct {
	boardRepository  domainBoard.BoardRepository
	columnRepository domainColumn.ColumnRepository
	taskRepository   domainTask.TaskRepository
}

func NewUseCase(
	boardRepository domainBoard.BoardRepository,
	columnRepository domainColumn.ColumnRepository,
	taskRepository domainTask.TaskRepository,
) *UseCase {
	return &UseCase{
		boardRepository:  boardRepository,
		columnRepository: columnRepository,
		taskRepository:   taskRepository,
	}
}

func (u *UseCase) GetBoard(ctx context.Context, id int) (domainCommon.Board, error) {
	board, err := u.boardRepository.GetById(ctx, id)
	if err != nil {
		return domainCommon.Board{}, err
	}

	columns, err := u.getColumns(ctx, board.Id)
	if err != nil {
		return domainCommon.Board{}, err
	}

	tasks, err := u.getTasks(ctx, columns)
	if err != nil {
		return domainCommon.Board{}, err
	}

	commonColumns := make([]domainCommon.Column, 0, len(columns))

	for _, column := range columns {
		var commonTasks []domainCommon.Task

		for _, task := range tasks {
			if task.ColumnId != column.Id {
				continue
			}

			commonTasks = append(commonTasks, domainCommon.Task{
				Id:          task.Id,
				Description: task.Description,
				Assignee:    task.Assignee,
				Position:    task.Position,
			})
		}

		commonColumns = append(commonColumns, domainCommon.Column{
			Id:       column.Id,
			Name:     column.Name,
			Position: column.Position,
			Tasks:    commonTasks,
		})
	}

	commonBoard := domainCommon.Board{
		Id:          board.Id,
		Name:        board.Name,
		Description: board.Description,
		OwnerId:     board.OwnerId,
		Columns:     commonColumns,
	}

	return commonBoard, err
}

func (u *UseCase) getColumns(ctx context.Context, boardId int) ([]domainColumn.Column, error) {
	columnFilter := domainColumn.NewGetAllFilter(boardId)
	columnSort := domainColumn.NewGetAllSort("asc")

	return u.columnRepository.GetAll(ctx, columnFilter, columnSort)
}

func (u *UseCase) getTasks(ctx context.Context, columns []domainColumn.Column) ([]domainTask.Task, error) {
	columnIds := make([]int, 0, len(columns))

	for _, column := range columns {
		columnIds = append(columnIds, column.Id)
	}

	taskFilter := domainTask.NewGetAllFilter(columnIds)
	taskSort := domainTask.NewGetAllSort("asc", "asc")

	return u.taskRepository.GetAll(ctx, taskFilter, taskSort)
}
