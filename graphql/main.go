package main

import (
	"context"
	"example-graphql-grpc/graphql/graph/generated"
	"example-graphql-grpc/graphql/resolver"
	"log"
	"net/http"

	"example-graphql-grpc/graphql/adapter/env"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	env := env.NewEnv()
	ctx := context.Background()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver.NewResolver(ctx, env),
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+env.GetGraphqlServerPort(), nil))
}
