package v1

import (
	"context"
	"net/http"

	"github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/create"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	"github.com/rs/zerolog/log"
)

func (a *Api) BoardCreate(ctx context.Context, req *gatewayV1.CreateBoardRequestBody) (gatewayV1.BoardCreateRes, error) {
	createBoardDTO := create.CreateBoardDTO{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		OwnerId:     req.GetOwnerID(),
	}

	boardId, err := a.boardCreateUseCase.Create(ctx, createBoardDTO)
	if err != nil {
		log.Error().Err(err).Msg("create board failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Create board failed",
		}, nil
	}

	return &gatewayV1.CreateBoardResponse{
		BoardID: boardId,
	}, nil
}

func (a *Api) BoardGetAll(ctx context.Context) (gatewayV1.BoardGetAllRes, error) {
	boards, err := a.boardGetAllUseCase.GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("get all board failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Get all boards failed",
		}, nil
	}

	apiBoards := make(gatewayV1.GetAllBoardResponse, 0, len(boards))

	for _, board := range boards {
		apiBoards = append(apiBoards, GetAllBoardDTOToTransport(board))
	}

	return &apiBoards, nil
}

func (a *Api) BoardGet(ctx context.Context, params gatewayV1.BoardGetParams) (gatewayV1.BoardGetRes, error) {
	board, err := a.boardGetBoardUseCase.Get(ctx, params.ID)
	if err != nil {
		log.Error().Err(err).Msg("get board failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Get board failed",
		}, nil
	}

	return GetBoardDTOToTransport(board), nil
}
