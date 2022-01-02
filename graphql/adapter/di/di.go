package di

import (
	userserviceaccessor "example-graphql-grpc/graphql/adapter/domainApi/userServiceAccessor"
	"example-graphql-grpc/graphql/shared"
	"reflect"

	"github.com/softia-inc/dject"
)

var Container dject.Container

func CreateContainer(
	env shared.Env,
) (dject.Container, error) {
	container := dject.NewContainer()
	Container = container
	// shared
	envOpt := dject.RegisterOptions{Interfaces: []reflect.Type{reflect.TypeOf((*shared.Env)(nil)).Elem()}}
	if err := container.Register(env, envOpt); err != nil {
		return nil, err
	}
	// shared
	if err := container.Register(userserviceaccessor.NewUserServiceAccessor); err != nil {
		return nil, err
	}
	return container, nil
}
