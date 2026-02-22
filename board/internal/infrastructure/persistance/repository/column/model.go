package column

import (
	"time"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
)

type Column struct {
	Id int `db:"id"`

	Name string `db:"name"`

	Position float64 `db:"position"`

	BoardId int `db:"board_id"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt *time.Time `db:"updated_at"`
}

func ConventModelToDomain(model Column) domainColumn.Column {
	return domainColumn.Column{
		Id:       model.Id,
		Name:     model.Name,
		Position: model.Position,
		BoardId:  model.BoardId,
	}
}
