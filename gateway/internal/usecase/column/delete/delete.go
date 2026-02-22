package delete

import (
	"context"

	"github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/column"
)

type UseCase struct {
	columnClient *column.Client
}

func NewUseCase(columnClient *column.Client) *UseCase {
	return &UseCase{
		columnClient: columnClient,
	}
}

func (u *UseCase) Delete(ctx context.Context, id int) error {
	deleteColumnRequest := column.DeleteColumnRequest{Id: id}

	err := u.columnClient.Delete(ctx, deleteColumnRequest)
	if err != nil {
		return err
	}

	return nil
}
