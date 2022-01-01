package rpc

import (
	definition "github.com/tkyatg/example-grpc-definition"
	"google.golang.org/grpc/reflection"
)

func (t *server) register(
	authCommandServiceServer definition.AuthCommandServiceServer,
) {
	definition.RegisterAuthCommandServiceServer(t.rpc, authCommandServiceServer)
	reflection.Register(t.rpc)
}
