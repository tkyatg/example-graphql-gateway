package resolver

import (
	"context"
	"example-graphql-grpc/graphql/adapter/domainapi/authserviceaccessor"
	"example-graphql-grpc/graphql/graph/generated"
	"example-graphql-grpc/graphql/shared"

	definition "github.com/tkyatg/example-grpc-definition"
	"google.golang.org/grpc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authCommandAccessor authserviceaccessor.AuthCommandServiceAccessor
}

func NewResolver(ctx context.Context, env shared.Env) generated.ResolverRoot {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// domain service
	conn, err := grpc.DialContext(ctx, env.GetDomainApiServerName()+":"+env.GetDomainApiPort(), opts...)
	if err != nil {
		panic(err)
	}
	// auth service
	authCommandClient := definition.NewAuthCommandServiceClient(conn)
	authCommandAccessor := authserviceaccessor.NewAuthServiceAccessor(authCommandClient)

	return &Resolver{
		authCommandAccessor,
	}
}
