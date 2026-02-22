package column

import (
	"context"

	columnConverter "github.com/poymanov/codemania-task-board/board/internal/infrastructure/converter/column"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetAll(ctx context.Context, req *boardV1.ColumnServiceGetAllRequest) (*boardV1.ColumnServiceGetAllResponse, error) {
	filter, sort := columnConverter.GetAllRequestToDomain(req)

	columns, err := s.columnGetAllUseCase.GetAll(ctx, filter, sort)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to get all boards")
		return nil, status.Error(codes.Internal, "failed to get all boards")
	}

	responseColumns := make([]*boardV1.Column, 0, len(columns))

	for _, column := range columns {
		responseColumns = append(responseColumns, columnConverter.DomainToTransport(column))
	}

	return &boardV1.ColumnServiceGetAllResponse{
		Columns: responseColumns,
	}, nil
}
