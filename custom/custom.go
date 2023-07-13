package custom

type (
	// PrefixBuilder is used to get custom prefix
	PrefixBuilder func() string
	CustomPrefix  struct {
		Builder PrefixBuilder
	}
)

// CreatePrefix create new CustomPrefix instance
func CreatePrefix(fn PrefixBuilder) *CustomPrefix {
	return &CustomPrefix{
		Builder: fn,
	}
}
