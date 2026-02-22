package board

import (
	"context"
	"errors"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	commonConverter "github.com/poymanov/codemania-task-board/board/internal/infrastructure/converter/common"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BoardService) GetBoard(ctx context.Context, req *boardV1.BoardServiceGetBoardRequest) (*boardV1.BoardServiceGetBoardResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	board, err := s.boardGetBoardUseCase.GetBoard(ctx, int(req.GetId()))
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to get board")

		if errors.Is(err, domainBoard.ErrBoardNotExists) {
			return nil, status.Errorf(codes.NotFound, "failed to get board: %v", err)
		}

		return nil, status.Error(codes.Internal, "failed to get board")
	}

	return commonConverter.BoardDomainToTransport(board), nil
}
