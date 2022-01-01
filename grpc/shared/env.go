package shared

type (
	Env interface {
		GetDBHost() string
		GetDBPort() string
		GetDBUser() string
		GetDBPassword() string
		GetDBName() string
		GetApiServerPort() string
		GetSignKey() string
	}
)
