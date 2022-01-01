package domain

import (
	"errors"

	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

type (
	authDataAccessor struct {
		dba shared.Sql
	}
)

func NewAuthAccessor(dba shared.Sql) AuthDataAccessor {
	return &authDataAccessor{
		dba,
	}
}
func (t *authDataAccessor) getUserByEmail(req *getUserByEmailRequest) (*getUserByEmailResult, error) {
	sql := `
select user_uuid
     , encripted_password
  from users
 where email = :email
;
`
	args := struct {
		Email string `db:"email"`
	}{
		Email: req.email,
	}
	rslt := struct {
		UserUUID          string `db:"user_uuid"`
		EncryptedPassword string `db:"encripted_password"`
	}{}
	rows, err := t.dba.Query(sql, args)
	if err != nil {
		return nil, errors.New(string(shared.SqlExecuteFail))
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&rslt); err != nil {
			return nil, err
		}
	}
	return &getUserByEmailResult{
		userUUID:          rslt.UserUUID,
		encryptedPassword: rslt.EncryptedPassword,
	}, nil
}
