package env

import (
	"example-graphql-grpc/graphql/shared"
	"os"
)

type (
	environment struct {
		graphqlServerPort   string
		domainApiServerName string
		domainApiServerPort string
	}
)

// NewEnv はコンストラクタです
func NewEnv() shared.Env {
	graphqlServerPort := os.Getenv("GRAPHQL_SERVICE_PORT")
	domainApiServerName := os.Getenv("DOMAIN_API_SERVICE_NAME")
	domainApiServerPort := os.Getenv("DOMAIN_API_SERVICE_PORT")
	return &environment{
		graphqlServerPort,
		domainApiServerName,
		domainApiServerPort,
	}
}

func (t *environment) GetGraphqlServerPort() string {
	return t.graphqlServerPort
}
func (t *environment) GetDomainApiServerName() string {
	return t.domainApiServerName
}
func (t *environment) GetDomainApiPort() string {
	return t.domainApiServerPort
}
