package pprint

type (
	DefaultPrefix struct {
		Message string
	}
)

func (p DefaultPrefix) Get() string {
	return p.Message
}

func (p DefaultPrefix) RegisterWord(words ...string) {}

func (p DefaultPrefix) SetFormat(format string) {}

func New(word string) *DefaultPrefix {
	return &DefaultPrefix{
		Message: word,
	}
}
