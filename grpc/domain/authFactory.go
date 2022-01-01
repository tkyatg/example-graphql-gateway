package domain

type (
	authFactory struct{}
	AuthFactory interface {
		CreateAuthorizationAttributes(email string, password string) (*authorizationAttributes, error)
	}
	authorizationAttributes struct {
		email    Email
		password Password
	}
)

func NewAuthFactory() AuthFactory {
	return &authFactory{}
}

func (t *authFactory) CreateAuthorizationAttributes(email string, password string) (*authorizationAttributes, error) {
	eml, err := newEmail(email)
	if err != nil {
		return nil, err
	}
	return &authorizationAttributes{
		email:    eml,
		password: Password(password),
	}, nil
}
