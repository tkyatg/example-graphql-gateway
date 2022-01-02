package rpc

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/tkyatg/example-graphql-grpc/grpc/adapter/env"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

type (
	authContext struct {
		ctx      context.Context
		authInfo *shared.AuthInfo
	}
	VerifyJwtTokenResponse struct {
		UserUUID string
	}
)

func NewAuthContext(ctx context.Context, authInfo *shared.AuthInfo) shared.AuthContext {
	return &authContext{ctx, authInfo}
}

func (t *authContext) GetContext() context.Context {
	return t.ctx
}

func (t *authContext) GetAuthInfo() *shared.AuthInfo {
	return t.authInfo
}

func verifyJwtToken(signedString string) (*VerifyJwtTokenResponse, error) {
	env := env.NewEnv()
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(string(shared.TokenUnexpected))
		}
		return []byte(env.GetSignKey()), nil
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

	return &VerifyJwtTokenResponse{
		UserUUID: userUUID,
	}, nil
}
