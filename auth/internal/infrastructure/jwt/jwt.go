package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	domainUser "github.com/poymanov/codemania-task-board/auth/internal/domain/user"
)

type JWT interface {
	GenerateAccessToken(user domainUser.User) (string, error)
	ValidateAccessToken(tokenString string) (domainUser.AuthClaims, error)
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
		"email":    user.Email,
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

func (j *JWTService) ValidateAccessToken(tokenString string) (domainUser.AuthClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return domainUser.AuthClaims{}, errors.New("failed to parse token")
		}
		return []byte(j.AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return domainUser.AuthClaims{}, fmt.Errorf("token is invalid: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return domainUser.AuthClaims{}, errors.New("failed to get claims from token")
	}

	if tokenType, ok := claims["type"].(string); !ok || tokenType != "access" {
		return domainUser.AuthClaims{}, errors.New("wrong token type")
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return domainUser.AuthClaims{}, errors.New("failed to get user_id")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return domainUser.AuthClaims{}, errors.New("failed to get email")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return domainUser.AuthClaims{}, errors.New("failed to get username")
	}

	return domainUser.AuthClaims{
		UserId:   int(userId),
		Email:    email,
		Username: username,
	}, nil
}
