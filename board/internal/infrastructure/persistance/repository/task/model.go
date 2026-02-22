package task

import (
	"time"

	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

type Task struct {
	Id int `db:"id"`

	Title string `db:"title"`

	Description string `db:"description"`

	Assignee string `db:"assignee"`

	Position float64 `db:"position"`

	ColumnId int `db:"column_id"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt *time.Time `db:"updated_at"`
}

func ConventModelToDomain(model Task) domainTask.Task {
	return domainTask.Task{
		Id:          model.Id,
		Title:       model.Title,
		Description: model.Description,
		Assignee:    model.Assignee,
		ColumnId:    model.ColumnId,
		Position:    model.Position,
	}
}
