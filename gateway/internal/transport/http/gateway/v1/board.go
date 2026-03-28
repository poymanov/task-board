package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/poymanov/codemania-task-board/gateway/internal/infrastructure/security"
	"github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/create"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	"github.com/rs/zerolog/log"
)

func (a *Api) BoardCreate(ctx context.Context, req *gatewayV1.CreateBoardRequestBody) (gatewayV1.BoardCreateRes, error) {
	userId, ok := security.GetUserID(ctx)

	if !ok {
		log.Error().Msg("Failed to get user id from context")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Create board failed",
		}, nil
	}

	createBoardDTO := create.CreateBoardDTO{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		OwnerId:     userId,
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
	reqStart := time.Now()
	board, err := a.boardGetBoardUseCase.Get(ctx, params.ID)
	status := http.StatusOK

	defer func() {
		a.httpMetrics.RequestsTotal.
			WithLabelValues("/api/v1/boards/{id}", http.MethodGet, strconv.Itoa(status)).
			Inc()

		a.httpMetrics.RequestDuration.
			WithLabelValues("/api/v1/boards/{id}", http.MethodGet).
			Observe(time.Since(reqStart).Seconds())
	}()

	if err != nil {
		status = http.StatusBadRequest

		log.Error().Err(err).Msg("get board failed")
		return &gatewayV1.BadRequestError{
			Code:    status,
			Message: "Get board failed",
		}, nil
	}

	log.Info().Int("id", params.ID).Msg("get board succeed")

	return GetBoardDTOToTransport(board), nil
}
