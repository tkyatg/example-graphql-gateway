package authserviceaccessor

import (
	"context"

	definition "github.com/tkyatg/example-grpc-definition"
)

type (
	authCommandServiceAccessor struct {
		client definition.AuthCommandServiceClient
	}
	AuthCommandServiceAccessor interface {
		Authorization(ctx context.Context, req *AuthorizationRequest) (*AuthorizationResponse, error)
	}
	AuthorizationRequest struct {
		Email    string
		Password string
	}
	AuthorizationResponse struct {
		JwtToken string
	}
)

func NewAuthServiceAccessor(client definition.AuthCommandServiceClient) AuthCommandServiceAccessor {
	return &authCommandServiceAccessor{
		client,
	}
}

func (r *authCommandServiceAccessor) Authorization(ctx context.Context, req *AuthorizationRequest) (*AuthorizationResponse, error) {
	res, err := r.client.Authorization(ctx, &definition.AuthorizationRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &AuthorizationResponse{
		JwtToken: res.JwtToken,
	}, nil
}
