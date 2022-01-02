package shared

import "context"

type (
	ContextKey  string
	TokenType   string
	ClaimKey    string
	AuthContext interface {
		GetAuthInfo() *AuthInfo
		GetContext() context.Context
	}
	AuthInfo struct {
		AuthenticationInfo
	}
	AuthenticationInfo struct {
		UserUUID string
	}
)

const (
	HeaderContextKey       ContextKey = "HeaderContextKey"
	ContainerContextKey    ContextKey = "ContainerContextKey"
	AuthorizationBearerKey ContextKey = "bearer"
	UserUUIDKey            ClaimKey   = "UserUUID"
	TokenTypeKey           ClaimKey   = "TokenType"
	UserAuthenticatedToken TokenType  = "UserAuthenticated"
)
