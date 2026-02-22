package v1

import (
	"context"
	"net/http"

	columnCreateUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/create"
	columnUpdatePositionUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/update_position"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	"github.com/rs/zerolog/log"
)

func (a *Api) ColumnCreate(ctx context.Context, req *gatewayV1.CreateColumnRequestBody, params gatewayV1.ColumnCreateParams) (gatewayV1.ColumnCreateRes, error) {
	createColumnDTO := columnCreateUseCase.CreateColumnDTO{
		Name:    req.GetName(),
		BoardId: params.ID,
	}

	columnId, err := a.columnCreateUseCase.Create(ctx, createColumnDTO)
	if err != nil {
		log.Error().Err(err).Msg("create column failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Create column failed",
		}, nil
	}

	return &gatewayV1.CreateColumnResponse{
		ColumnID: columnId,
	}, nil
}

func (a *Api) ColumnDelete(ctx context.Context, params gatewayV1.ColumnDeleteParams) (gatewayV1.ColumnDeleteRes, error) {
	err := a.columnDeleteUseCase.Delete(ctx, params.ColumnId)
	if err != nil {
		log.Error().Err(err).Msg("create column failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Delete column failed",
		}, nil
	}

	return &gatewayV1.ColumnDeleteNoContent{}, nil
}

func (a *Api) ColumnUpdatePosition(ctx context.Context, req *gatewayV1.ColumnUpdatePositionRequestBody, params gatewayV1.ColumnUpdatePositionParams) (gatewayV1.ColumnUpdatePositionRes, error) {
	updatePositionColumnDTO := columnUpdatePositionUseCase.UpdatePositionColumnDTO{
		Id:            params.ColumnId,
		LeftPosition:  req.LeftPosition,
		RightPosition: req.RightPosition,
	}

	err := a.columnUpdatePositionUseCase.UpdatePosition(ctx, updatePositionColumnDTO)
	if err != nil {
		log.Error().Err(err).Any("request", req).Any("params", params).Msg("update position column failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Change column position failed",
		}, nil
	}

	return &gatewayV1.ColumnUpdatePositionNoContent{}, nil
}
