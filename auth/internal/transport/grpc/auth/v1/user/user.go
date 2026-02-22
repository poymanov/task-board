package user

import (
	registerUserUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/register"
	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
)

type Service struct {
	registerUserUseCase *registerUserUseCase.UseCase

	authV1.UnimplementedUserServiceServer
}

func NewService(registerUserUseCase *registerUserUseCase.UseCase) *Service {
	return &Service{
		registerUserUseCase: registerUserUseCase,
	}
}
