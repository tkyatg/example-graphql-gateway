package authcommand

import (
	"github.com/tkyatg/example-graphql-grpc/grpc/domain"
)

type (
	usecase struct {
		factory domain.AuthFactory
		repo    domain.AuthRepository
	}
	Usecase interface {
		authorization(req *authorizationRequest) (*authorizationResponse, error)
	}
	authorizationRequest struct {
		email    string
		password string
	}
	authorizationResponse struct {
		jwtToken string
	}
)

func NewUsecase(factory domain.AuthFactory, repo domain.AuthRepository) Usecase {
	return &usecase{
		factory,
		repo,
	}
}

func (t *usecase) authorization(req *authorizationRequest) (*authorizationResponse, error) {
	attr, err := t.factory.CreateAuthorizationAttributes(req.email, req.password)
	if err != nil {
		return nil, err
	}
	res, err := t.repo.Authorization(attr)
	if err != nil {
		return nil, err
	}
	return &authorizationResponse{
		jwtToken: res.JwtToken,
	}, nil
}
