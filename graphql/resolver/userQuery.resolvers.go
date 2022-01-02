package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	userserviceaccessor "example-graphql-grpc/graphql/adapter/domainApi/userServiceAccessor"
	"example-graphql-grpc/graphql/graph/generated/model"
	"example-graphql-grpc/graphql/shared"

	"github.com/softia-inc/dject"
)

func (r *queryResolver) GetLoginUser(ctx context.Context) (*model.GetLoginUserResponse, error) {
	container := ctx.Value(shared.ContainerContextKey).(dject.Container)
	res := &model.GetLoginUserResponse{}
	if err := container.Invoke(func(accessor userserviceaccessor.UserQueryServiceAccessor) error {
		rslt, err := accessor.GetLoginUser(ctx, &userserviceaccessor.GetLoginUserRequest{})
		if err != nil {
			return err
		}
		res = &model.GetLoginUserResponse{
			UserUUID: rslt.UserUUID,
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}
