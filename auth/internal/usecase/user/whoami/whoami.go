package whoami

import (
	"github.com/poymanov/codemania-task-board/auth/internal/infrastructure/jwt"
)

type UseCase struct {
	jwtService jwt.JWT
}

func NewUseCase(jwtService jwt.JWT) *UseCase {
	return &UseCase{
		jwtService: jwtService,
	}
}

func (u *UseCase) Whoami(accessToken string) (WhoamiDTO, error) {
	claims, err := u.jwtService.ValidateAccessToken(accessToken)
	if err != nil {
		return WhoamiDTO{}, err
	}

	return WhoamiDTO{
		UserId:   claims.UserId,
		Email:    claims.Email,
		Username: claims.Username,
	}, nil
}
