package user

import (
	"context"

	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Whoami(ctx context.Context, req *authV1.UserServiceWhoamiRequest) (*authV1.UserServiceWhoamiResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	dto, err := s.whoamiUseCase.Whoami(req.AccessToken)
	if err != nil {
		log.Error().Err(err).Msg("failed to check user")

		return nil, status.Error(codes.Unauthenticated, domainUser.ErrInvalidToken.Error())
	}

	return &authV1.UserServiceWhoamiResponse{
		UserId:   int64(dto.UserId),
		Email:    dto.Email,
		Username: dto.Username,
	}, nil
}
