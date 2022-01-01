package hash

import (
	"errors"
	"fmt"

	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
	"golang.org/x/crypto/bcrypt"
)

type (
	hash struct {
		env shared.Env
	}
)

func NewHash(env shared.Env) shared.Hash {
	return &hash{
		env,
	}
}

func (t *hash) Generate(attr string) (shared.HashedString, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(attr), 10)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(string(shared.HashFail))
	}
	return shared.HashedString(res), nil
}

func (t *hash) IsSameString(hashedString string, inputString string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(inputString)); err != nil {
		return false
	}
	return true
}
