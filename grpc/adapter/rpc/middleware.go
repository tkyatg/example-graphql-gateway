package rpc

import (
	"context"
	"errors"

	"github.com/softia-inc/dject"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func unaryServerInterceptor(container dject.Container) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			err1 := errors.New("request header の取得に失敗しました")
			err = status.Error(codes.Internal, err1.Error())
			return
		}

		childContainer := container.CreateChildContainer()

		ctx = context.WithValue(ctx, shared.HeaderContextKey, md)
		ctx = context.WithValue(ctx, shared.ContainerContextKey, childContainer)
		return handler(ctx, req)

	}
}
