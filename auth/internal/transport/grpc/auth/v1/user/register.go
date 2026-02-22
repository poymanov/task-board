package user

import (
	"context"
	"errors"

	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	registerUserUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/register"
	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Register(ctx context.Context, req *authV1.UserServiceRegisterRequest) (*authV1.UserServiceRegisterResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	dto := registerUserUseCase.RegisterUserDTO{
		Password: req.Password,
		Email:    req.Email,
		Username: req.Username,
	}

	err := s.registerUserUseCase.Register(ctx, dto)
	if err != nil {
		errMessage := "failed to register user"

		log.Error().Err(err).Msg(errMessage)

		if errors.Is(err, domainUser.ErrUserAlreadyExists) {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.Internal, errMessage)
	}

	return &authV1.UserServiceRegisterResponse{}, nil
}
