package rpc

import (
	definition "github.com/tkyatg/example-grpc-definition"
	"google.golang.org/grpc/reflection"
)

func (t *server) register(
	authCommandServiceServer definition.AuthCommandServiceServer,
	userQueryServiceServer definition.UserQueryServiceServer,
) {
	definition.RegisterAuthCommandServiceServer(t.rpc, authCommandServiceServer)
	definition.RegisterUserQueryServiceServer(t.rpc, userQueryServiceServer)
	reflection.Register(t.rpc)
}
