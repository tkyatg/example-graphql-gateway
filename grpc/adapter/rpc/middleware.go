package rpc

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/softia-inc/dject"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type (
	needAuthenticationMethodName string
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
			return nil, status.Error(codes.Internal, errors.New("request header の取得に失敗しました").Error())
		}
		childContainer := container.CreateChildContainer()
		// 認証確認
		if needAuthenticationMethods(info.FullMethod) {
			signedString := strings.Replace(md["authorization"][0], "Bearer ", "", 1)
			if signedString == "" {
				return nil, status.Error(codes.Unauthenticated, errors.New("authorization token が空です").Error())
			}
			verifyJwtTokenRes, err := verifyJwtToken(signedString)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}
			authInfo := &shared.AuthInfo{
				AuthenticationInfo: shared.AuthenticationInfo{
					UserUUID: verifyJwtTokenRes.UserUUID,
				},
			}
			authOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*shared.AuthContext)(nil)).Elem()}}
			if err := childContainer.Register(NewAuthContext(ctx, authInfo), authOpt); err != nil {
				return nil, err
			}
		}

		ctx = context.WithValue(ctx, shared.HeaderContextKey, md)
		ctx = context.WithValue(ctx, shared.ContainerContextKey, childContainer)
		return handler(ctx, req)

	}
}

func needAuthenticationMethods(methodName string) bool {
	methods := []needAuthenticationMethodName{
		"/example.UserQueryService/GetLoginUser",
	}
	for _, needAuthenticationMethodName := range methods {
		if string(needAuthenticationMethodName) == methodName {
			return true
		}
	}
	return false
}
