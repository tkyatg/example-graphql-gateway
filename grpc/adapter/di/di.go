package di

import (
	"reflect"

	"github.com/softia-inc/dject"
	"github.com/tkyatg/example-graphql-grpc/grpc/adapter/hash"
	authcommand "github.com/tkyatg/example-graphql-grpc/grpc/commands/authCommand"
	"github.com/tkyatg/example-graphql-grpc/grpc/domain"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

var Container dject.Container

func CreateContainer(
	env shared.Env,
	dba shared.Sql,
) (dject.Container, error) {
	container := dject.NewContainer()
	Container = container
	// shared
	envOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*shared.Env)(nil)).Elem()}}
	if err := container.Register(env, envOpt); err != nil {
		return nil, err
	}
	dbOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*shared.Sql)(nil)).Elem()}}
	if err := container.Register(dba, dbOpt); err != nil {
		return nil, err
	}
	hashOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*shared.Hash)(nil)).Elem()}}
	if err := container.Register(hash.NewHash, hashOpt); err != nil {
		return nil, err
	}

	// command
	// hostcommand
	if err := container.Register(authcommand.NewServer); err != nil {
		return nil, err
	}
	if err := container.Register(authcommand.NewUsecase); err != nil {
		return nil, err
	}

	// domain
	if err := container.Register(domain.NewAuthFactory); err != nil {
		return nil, err
	}
	if err := container.Register(domain.NewAuthRepository); err != nil {
		return nil, err
	}
	if err := container.Register(domain.NewAuthAccessor); err != nil {
		return nil, err
	}
	if err := container.Register(domain.NewTokenRepository); err != nil {
		return nil, err
	}

	return container, nil
}
