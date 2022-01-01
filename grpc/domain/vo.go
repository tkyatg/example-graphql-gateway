package domain

import (
	"errors"
	"regexp"

	"github.com/tkyatg/example-graphql-grpc/grpc/shared"
)

type (
	Email    string
	Password string
	JwtToken string
	Token    struct{}
)

var emailFormat = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`)

func newEmail(attr string) (Email, error) {
	if attr == "" {
		return "", errors.New(string(shared.EmailReuqired))
	}
	if !emailFormat.MatchString(attr) {
		return "", errors.New(string(shared.EmailReuqired))
	}
	return Email(attr), nil
}
