package shared

type (
	errorCode string
)

const (
	SqlExecuteFail   errorCode = "error-000001"
	EmailReuqired    errorCode = "error-000002"
	PasswordReuqired errorCode = "error-000003"
	PasswordTooShort errorCode = "error-000004"
	PasswordHashFail errorCode = "error-000005"
	HashFail         errorCode = "error-000006"
	PasswordFail     errorCode = "error-000007"
	TokenUnexpected  errorCode = "error-000008"
	TokenExpired     errorCode = "error-000009"
	TokenInvalid     errorCode = "error-000010"
	TokenNotFound    errorCode = "error-000011"
	ClaimNotFound    errorCode = "error-000012"
)
