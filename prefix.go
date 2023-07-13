package pprint

type (
	Prefix interface {
		// Get return registered prefix
		Get() string
	}
)
