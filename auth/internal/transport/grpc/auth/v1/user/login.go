package user

import (
	"context"
	"errors"

	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
	loginUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/login"
	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Login(ctx context.Context, req *authV1.UserServiceLoginRequest) (*authV1.UserServiceLoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	dto := loginUseCase.LoginDTO{Email: req.Email, Password: req.Password}

	accessToken, err := s.loginUseCase.Login(ctx, dto)
	if err != nil {
		errMessage := "failed to login"

		log.Error().Err(err).Msg(errMessage)

		if errors.Is(err, domainUser.ErrInvalidCredentials) {
			return nil, status.Error(codes.Internal, domainUser.ErrInvalidCredentials.Error())
		}

		return nil, status.Error(codes.Internal, errMessage)
	}

	return &authV1.UserServiceLoginResponse{AccessToken: accessToken}, nil
}
