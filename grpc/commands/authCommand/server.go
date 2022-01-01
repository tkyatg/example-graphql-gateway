package authcommand

import (
	"context"

	"github.com/softia-inc/dject"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
	definition "github.com/tkyatg/example-grpc-definition"
)

type server struct {
}

func NewServer() definition.AuthCommandServiceServer {
	return &server{}
}

func (t *server) Authorization(ctx context.Context, req *definition.AuthorizationRequest) (*definition.AuthorizationResponse, error) {
	container := ctx.Value(shared.ContainerContextKey).(dject.Container)
	res := &definition.AuthorizationResponse{}
	if err := container.Invoke(func(uc Usecase) error {
		rslt, err := uc.authorization(&authorizationRequest{
			email:    req.Email,
			password: req.Password,
		})
		if err != nil {
			return err
		}
		res = &definition.AuthorizationResponse{
			JwtToken: rslt.jwtToken,
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}
