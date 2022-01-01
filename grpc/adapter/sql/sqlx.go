package sql

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

type (
	SqlxConnectRequest struct {
		DbUser     string
		DbPassword string
		DbName     string
		DbPort     string
		DbHost     string
	}
	sql struct {
		dba *sqlx.DB
		bk  sqlx.DB
	}
)

// NewSqlxConnect func
func NewSqlxConnect(attr *SqlxConnectRequest) (shared.Sql, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", attr.DbHost, attr.DbUser, attr.DbName, attr.DbPassword)
	dba, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return nil, err
	}

	return &sql{
		dba: dba,
		bk:  *dba,
	}, nil

}
func (t *sql) Query(query string, args interface{}) (*sqlx.Rows, error) {
	return t.dba.NamedQuery(query, args)
}
func (t *sql) Exec(query string, args interface{}) error {
	if _, err := t.dba.NamedExec(query, args); err != nil {
		return err
	}
	return nil
}
