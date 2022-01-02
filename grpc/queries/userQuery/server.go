package userquery

import (
	"context"

	"github.com/softia-inc/dject"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
	definition "github.com/tkyatg/example-grpc-definition"
)

type server struct {
}

func NewServer() definition.UserQueryServiceServer {
	return &server{}
}

func (t *server) GetLoginUser(ctx context.Context, req *definition.GetLoginUserRequest) (*definition.GetLoginUserResponse, error) {
	container := ctx.Value(shared.ContainerContextKey).(dject.Container)
	res := &definition.GetLoginUserResponse{}
	if err := container.Invoke(func(atx shared.AuthContext) error {
		res = &definition.GetLoginUserResponse{
			UserUUID: atx.GetAuthInfo().UserUUID,
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}
