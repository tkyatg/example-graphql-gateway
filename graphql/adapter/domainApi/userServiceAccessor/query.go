package userserviceaccessor

import (
	"context"
	"example-graphql-grpc/graphql/shared"

	definition "github.com/tkyatg/example-grpc-definition"
	"google.golang.org/grpc"
)

type (
	userQueryServiceAccessor struct {
		env shared.Env
	}
	UserQueryServiceAccessor interface {
		GetLoginUser(ctx context.Context, req *GetLoginUserRequest) (*GetLoginUserResponse, error)
	}
	GetLoginUserRequest struct {
	}
	GetLoginUserResponse struct {
		UserUUID string
	}
)

func NewUserServiceAccessor(env shared.Env) UserQueryServiceAccessor {
	return &userQueryServiceAccessor{
		env,
	}
}

func (t *userQueryServiceAccessor) GetLoginUser(ctx context.Context, req *GetLoginUserRequest) (*GetLoginUserResponse, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// user service
	conn, err := grpc.DialContext(ctx, t.env.GetDomainApiServerName()+":"+t.env.GetDomainApiPort(), opts...)
	if err != nil {
		panic(err)
	}
	// auth service
	userQueryClient := definition.NewUserQueryServiceClient(conn)
	res, err := userQueryClient.GetLoginUser(ctx, &definition.GetLoginUserRequest{})
	if err != nil {
		return nil, err
	}
	return &GetLoginUserResponse{
		UserUUID: res.UserUUID,
	}, nil
}
