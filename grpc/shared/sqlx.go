package shared

import "github.com/jmoiron/sqlx"

type (
	Sql interface {
		Query(query string, args interface{}) (*sqlx.Rows, error)
		Exec(query string, args interface{}) error
	}
)
