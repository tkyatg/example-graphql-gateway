package env

import (
	"os"

	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

type (
	environment struct {
		dbHost        string
		dbPort        string
		dbUser        string
		dbPassword    string
		dbName        string
		apiServerPort string
		signKey       string
	}
)

func NewEnv() shared.Env {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	apiServerPort := os.Getenv("PORT")
	signKey := os.Getenv("SIGN_KEY")
	return &environment{
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
		apiServerPort,
		signKey,
	}
}

func (t *environment) GetDBHost() string {
	return t.dbHost
}
func (t *environment) GetDBPort() string {
	return t.dbPort
}
func (t *environment) GetDBUser() string {
	return t.dbUser
}
func (t *environment) GetDBPassword() string {
	return t.dbPassword
}
func (t *environment) GetDBName() string {
	return t.dbName
}
func (t *environment) GetApiServerPort() string {
	return t.apiServerPort
}
func (t *environment) GetSignKey() string {
	return t.signKey
}
