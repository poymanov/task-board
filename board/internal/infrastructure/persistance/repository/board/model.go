package board

import (
	"time"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

type Board struct {
	Id int `db:"id"`

	Name string `db:"name"`

	Description string `db:"description"`

	OwnerId int `db:"owner_id"`

	CreatedAt time.Time `db:"created_at"`

	UpdatedAt *time.Time `db:"updated_at"`
}

func ConvertModelToDomain(board Board) domainBoard.Board {
	return domainBoard.Board{
		Id:          board.Id,
		Name:        board.Name,
		Description: board.Description,
		OwnerId:     board.OwnerId,
	}
}
