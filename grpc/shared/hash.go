package shared

type (
	HashedString string
	Hash         interface {
		Generate(attr string) (HashedString, error)
		IsSameString(hashedString string, inputString string) bool
	}
)
