package user

import (
	loginUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/login"
	registerUserUseCase "github.com/poymanov/codemania-task-board/auth/internal/usecase/user/register"
	authV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/auth/v1"
)

type Service struct {
	registerUserUseCase *registerUserUseCase.UseCase

	loginUseCase *loginUseCase.UseCase

	authV1.UnimplementedUserServiceServer
}

func NewService(registerUserUseCase *registerUserUseCase.UseCase, loginUseCase *loginUseCase.UseCase) *Service {
	return &Service{
		registerUserUseCase: registerUserUseCase,
		loginUseCase:        loginUseCase,
	}
}
