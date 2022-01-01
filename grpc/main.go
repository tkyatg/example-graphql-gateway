package main

import (
	"log"

	"github.com/tkyatg/example-graphql-grpc/grpc/adapter/di"
	"github.com/tkyatg/example-graphql-grpc/grpc/adapter/env"
	"github.com/tkyatg/example-graphql-grpc/grpc/adapter/rpc"
	"github.com/tkyatg/example-graphql-grpc/grpc/adapter/sql"
)

func main() {
	env := env.NewEnv()
	dba, err := sql.NewSqlxConnect(&sql.SqlxConnectRequest{
		DbUser:     env.GetDBUser(),
		DbPassword: env.GetDBPassword(),
		DbName:     env.GetDBName(),
		DbPort:     env.GetDBPort(),
		DbHost:     env.GetDBHost(),
	})
	if err != nil {
		log.Fatal(err)
	}
	container, err := di.CreateContainer(env, dba)
	if err != nil {
		log.Fatal(err)
	}
	server := rpc.NewServer(env.GetApiServerPort(), container)
	server.Serve()
}
