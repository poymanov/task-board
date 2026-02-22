package column

import (
	"context"
	"errors"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	columnCreateUsecase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/create"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Create(ctx context.Context, req *boardV1.ColumnServiceCreateRequest) (*boardV1.ColumnServiceCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	newColumn := columnCreateUsecase.NewColumnDTO{Name: req.GetName(), BoardID: int(req.GetBoardId())}

	id, err := s.columnCreateUseCase.Create(ctx, newColumn)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to create column")

		if errors.Is(err, domainBoard.ErrBoardNotExists) {
			return nil, status.Errorf(codes.Internal, "failed to create column: %v", err)
		}

		return nil, status.Error(codes.Internal, "failed to create column")
	}

	return &boardV1.ColumnServiceCreateResponse{ColumnId: int64(id)}, nil
}
