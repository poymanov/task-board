package update_position

import (
	"context"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
)

type UseCase struct {
	columnRepository domainColumn.ColumnRepository
}

func NewUseCase(columnRepository domainColumn.ColumnRepository) *UseCase {
	return &UseCase{
		columnRepository: columnRepository,
	}
}

func (u *UseCase) UpdatePosition(ctx context.Context, id int, updatePositionDTO UpdatePositionDTO) error {
	exist, err := u.columnRepository.IsExistsById(ctx, id)
	if err != nil {
		return err
	}

	if !exist {
		return domainColumn.ErrColumnNotExists
	}

	newPosition := (updatePositionDTO.LeftPosition + updatePositionDTO.RightPosition) / 2

	err = u.columnRepository.UpdatePosition(ctx, id, newPosition)
	if err != nil {
		return err
	}

	return nil
}
