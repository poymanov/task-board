package v1

import (
	"context"
	"net/http"

	authRegisterUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/auth/register"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	"github.com/rs/zerolog/log"
)

func (a *Api) AuthRegister(ctx context.Context, req *gatewayV1.RegisterRequestBody) (gatewayV1.AuthRegisterRes, error) {
	dto := authRegisterUseCase.RegisterDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Username: req.GetUsername(),
	}

	err := a.authRegisterUseCase.Register(ctx, dto)
	if err != nil {
		errMessage := "register failed"

		log.Error().Err(err).Msg(errMessage)
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: errMessage,
		}, nil
	}

	return &gatewayV1.AuthRegisterCreated{}, nil
}
