package main

import (
	"context"
	"example-graphql-grpc/graphql/graph/generated"
	"example-graphql-grpc/graphql/resolver"
	"net/http"

	"example-graphql-grpc/graphql/adapter/di"
	"example-graphql-grpc/graphql/adapter/env"
	authmiddleware "example-graphql-grpc/graphql/middleware/authMiddleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func main() {
	env := env.NewEnv()
	router := chi.NewRouter()
	di, err := di.CreateContainer(env)
	if err != nil {
		panic(err)
	}
	router.Use(authmiddleware.Middleware(di))

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver.NewResolver(context.Background(), env),
			},
		),
	)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	if err := http.ListenAndServe(":"+env.GetGraphqlServerPort(), router); err != nil {
		panic(err)
	}
}
