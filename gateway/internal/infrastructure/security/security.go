package security

import (
	"context"
	"errors"

	authWhoamiUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/auth/whoami"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrFailedToCheckAuth = errors.New("failed to check auth")
)

type userIDKey struct{}

type SecurityHandler struct {
	authWhoamiUseCase *authWhoamiUseCase.UseCase
}

func NewSecurityHandler(authWhoamiUseCase *authWhoamiUseCase.UseCase) *SecurityHandler {
	return &SecurityHandler{authWhoamiUseCase: authWhoamiUseCase}
}

func (s *SecurityHandler) HandleBearerAuth(
	ctx context.Context,
	operationName gatewayV1.OperationName,
	t gatewayV1.BearerAuth,
) (context.Context, error) {
	dto := authWhoamiUseCase.WhoamiDTO{AccessToken: t.Token}

	user, err := s.authWhoamiUseCase.Whoami(ctx, dto)
	if err != nil {
		log.Error().Str("operation", operationName).Err(err).Msg("handle bearer auth failed")

		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.Unauthenticated:
				return ctx, ErrInvalidToken
			default:
				return ctx, ErrFailedToCheckAuth
			}
		}
		return ctx, ErrFailedToCheckAuth
	}

	ctx = context.WithValue(ctx, userIDKey{}, user.UserId)

	return ctx, nil
}

func GetUserID(ctx context.Context) (int, bool) {
	v := ctx.Value(userIDKey{})
	id, ok := v.(int)
	return id, ok
}
