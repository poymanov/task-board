package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
)

type JWT interface {
	GenerateAccessToken(user domainUser.User) (string, error)
}

type JWTService struct {
	AccessTokenTTL    time.Duration
	AccessTokenSecret string
}

func NewJWTService(accessTokenTTL time.Duration, accessTokenSecret string) *JWTService {
	return &JWTService{
		AccessTokenTTL:    accessTokenTTL,
		AccessTokenSecret: accessTokenSecret,
	}
}

func (j *JWTService) GenerateAccessToken(user domainUser.User) (string, error) {
	expiresAt := time.Now().Add(j.AccessTokenTTL)

	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Username,
		"email":    user.Username,
		"exp":      expiresAt.Unix(),
		"iat":      time.Now().Unix(),
		"type":     "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
