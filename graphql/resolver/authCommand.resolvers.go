package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example-graphql-grpc/graphql/adapter/domainapi/authserviceaccessor"
	"example-graphql-grpc/graphql/graph/generated/model"
)

func (r *mutationResolver) Authorize(ctx context.Context, input model.AuthorizeRequest) (*model.AuthorizeResponse, error) {
	res, err := r.accessor.Authorization(ctx, &authserviceaccessor.AuthorizationRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return &model.AuthorizeResponse{
		JwtToken: res.JwtToken,
	}, nil
}
