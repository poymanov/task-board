package v1

import (
	"context"
	"net/http"

	taskCreateUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/task/create"
	taskUpdatePositionUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/task/update_position"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	"github.com/rs/zerolog/log"
)

func (a *Api) TaskCreate(ctx context.Context, req *gatewayV1.TaskCreateRequestBody, params gatewayV1.TaskCreateParams) (gatewayV1.TaskCreateRes, error) {
	taskCreateDTO := taskCreateUseCase.TaskCreateDTO{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Assignee:    req.GetAssignee(),
		ColumnId:    params.ColumnId,
	}

	taskId, err := a.taskCreateUseCase.Create(ctx, taskCreateDTO)
	if err != nil {
		log.Error().Err(err).Msg("create task failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Create task failed",
		}, nil
	}

	return &gatewayV1.TaskCreateResponse{
		TaskID: taskId,
	}, nil
}

func (a *Api) TaskDelete(ctx context.Context, params gatewayV1.TaskDeleteParams) (gatewayV1.TaskDeleteRes, error) {
	err := a.taskDeleteUseCase.Delete(ctx, params.TaskId)
	if err != nil {
		log.Error().Err(err).Msg("delete task failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Delete task failed",
		}, nil
	}

	return &gatewayV1.TaskDeleteNoContent{}, nil
}

func (a *Api) TaskUpdatePosition(ctx context.Context, req *gatewayV1.TaskUpdatePositionRequestBody, params gatewayV1.TaskUpdatePositionParams) (gatewayV1.TaskUpdatePositionRes, error) {
	updatePositionTaskDTO := taskUpdatePositionUseCase.UpdatePositionColumnDTO{
		Id:            params.TaskId,
		LeftPosition:  req.LeftPosition,
		RightPosition: req.RightPosition,
	}

	err := a.taskUpdatePositionUseCase.UpdatePosition(ctx, updatePositionTaskDTO)
	if err != nil {
		log.Error().Err(err).Any("request", req).Any("params", params).Msg("update position task failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Update task position failed",
		}, nil
	}

	return &gatewayV1.TaskUpdatePositionNoContent{}, nil
}
