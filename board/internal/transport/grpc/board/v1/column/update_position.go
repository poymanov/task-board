package column

import (
	"context"
	"errors"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	columnConverter "github.com/poymanov/codemania-task-board/board/internal/infrastructure/converter/column"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) UpdatePosition(ctx context.Context, req *boardV1.ColumnServiceUpdatePositionRequest) (*boardV1.ColumnServiceUpdatePositionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	id := int(req.GetId())

	dto := columnConverter.UpdatePositionRequestToUseCaseDTO(req)

	err := s.columnUpdatePositionUseCase.UpdatePosition(ctx, id, dto)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to update column position")

		if errors.Is(err, domainColumn.ErrColumnNotExists) {
			return nil, status.Errorf(codes.Internal, "failed to update column position: %v", err)
		}

		return nil, status.Error(codes.Internal, "failed to update column position")
	}

	return &boardV1.ColumnServiceUpdatePositionResponse{}, nil
}
