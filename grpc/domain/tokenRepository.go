package domain

import (
	"errors"
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
		VerifyJwtToken(signedString string) (*VerifyJwtTokenResponse, error)
	}
	VerifyJwtTokenResponse struct {
		UserUUID  string
		TokenType string
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

func (t *tokenRepository) VerifyJwtToken(signedString string) (*VerifyJwtTokenResponse, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(string(shared.TokenUnexpected))
		}
		return []byte(t.env.GetSignKey()), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(string(shared.TokenExpired))
			} else {
				return nil, errors.New(string(shared.TokenInvalid))
			}
		} else {
			return nil, errors.New(string(shared.TokenInvalid))
		}
	}

	if token == nil {
		return nil, errors.New(string(shared.TokenNotFound))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(string(shared.ClaimNotFound))
	}
	userUUID := ""
	uUUID, ok := claims[string(shared.UserUUIDKey)]
	if ok {
		userUUID = uUUID.(string)
	}
	tokenType := ""
	tType, ok := claims[string(shared.TokenTypeKey)]
	if ok {
		tokenType = tType.(string)
	}

	return &VerifyJwtTokenResponse{
		UserUUID:  userUUID,
		TokenType: tokenType,
	}, nil
}
