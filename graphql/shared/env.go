package shared

type (
	// Env interface
	Env interface {
		GetGraphqlServerPort() string
		GetDomainApiServerName() string
		GetDomainApiPort() string
	}
)
