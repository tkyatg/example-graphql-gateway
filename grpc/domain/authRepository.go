package domain

import (
	"errors"

	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

type (
	authRepository struct {
		tokenRepo TokenRepository
		da        AuthDataAccessor
		hash      shared.Hash
	}
	AuthRepository interface {
		Authorization(attr *authorizationAttributes) (*AuthorizationResponse, error)
	}
	AuthDataAccessor interface {
		getUserByEmail(req *getUserByEmailRequest) (*getUserByEmailResult, error)
	}
	getUserByEmailRequest struct {
		email string
	}
	getUserByEmailResult struct {
		userUUID          string
		encryptedPassword string
	}
	AuthorizationResponse struct {
		JwtToken string
	}
)

func NewAuthRepository(tokenRepo TokenRepository, da AuthDataAccessor, hash shared.Hash) AuthRepository {
	return &authRepository{
		tokenRepo,
		da,
		hash,
	}
}
func (t *authRepository) Authorization(attr *authorizationAttributes) (*AuthorizationResponse, error) {
	res, err := t.da.getUserByEmail(&getUserByEmailRequest{
		email: string(attr.email),
	})
	if err != nil {
		return nil, err
	}
	if !t.hash.IsSameString(res.encryptedPassword, string(attr.password)) {
		return nil, errors.New(string(shared.PasswordFail))
	}
	jwtToken, err := t.tokenRepo.GenerateJwtToken(map[shared.ClaimKey]interface{}{
		shared.UserUUIDKey: res.userUUID,
	})
	if err != nil {
		return nil, err
	}
	return &AuthorizationResponse{
		JwtToken: string(jwtToken),
	}, nil
}
