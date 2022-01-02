package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

type (
	tokenRepository struct {
		env shared.Env
	}
	TokenRepository interface {
		GenerateJwtToken(attr map[shared.ClaimKey]interface{}) (JwtToken, error)
	}
)

func NewTokenRepository(env shared.Env) TokenRepository {
	return &tokenRepository{
		env,
	}
}

func (t *tokenRepository) GenerateJwtToken(attr map[shared.ClaimKey]interface{}) (JwtToken, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 24).Unix(),
	}
	for key, val := range attr {
		claims[string(key)] = val
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.env.GetSignKey()))
	if err != nil {
		return "", err
	}
	return JwtToken(tokenString), nil
}
