package authmiddleware

import (
	"context"
	"net/http"

	"example-graphql-grpc/graphql/shared"

	"github.com/softia-inc/dject"
	"google.golang.org/grpc/metadata"
)

var TokenKey = "JWT_TOKEN"

func Middleware(container dject.Container) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			md := metadata.New(map[string]string{"Authorization": r.Header.Get("Authorization")})
			ctx := metadata.NewOutgoingContext(r.Context(), md)
			ctx = context.WithValue(ctx, shared.ContainerContextKey, container)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
