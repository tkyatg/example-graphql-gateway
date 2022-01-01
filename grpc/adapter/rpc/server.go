package rpc

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/softia-inc/dject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type (
	server struct {
		port      string
		rpc       *grpc.Server
		container dject.Container
	}
	Server interface {
		Serve() error
	}
)

func NewServer(port string, container dject.Container) *server {
	s := &server{
		port:      port,
		container: container,
	}
	if err := s.setRPC(); err != nil {
		log.Fatal(err)
	}
	return s
}

func (t *server) Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", t.port))
	if err != nil {
		log.Fatal(err)
	}
	t.rpc.Serve(lis)
}

func (t *server) setRPC() error {
	t.rpc = grpc.NewServer(
		grpc.UnaryInterceptor(unaryServerInterceptor(t.container)),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:                  15 * time.Second,
				Timeout:               30 * time.Second,
				MaxConnectionAge:      90 * time.Second,
				MaxConnectionAgeGrace: 5 * time.Hour,
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             10 * time.Second,
				PermitWithoutStream: true,
			},
		),
	)

	return t.container.Invoke(t.register)
}
