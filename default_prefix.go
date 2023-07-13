package pprint

import "fmt"

type (
	DefaultPrefix struct {
		Message string
	}
)

func (p DefaultPrefix) Get() string {
	return fmt.Sprintf("[%s]", p.Message)
}

func (p DefaultPrefix) RegisterWord(words ...string) {}

func (p DefaultPrefix) SetFormat(format string) {}

// New return new DefaultPrefix instance
func New(word string) *DefaultPrefix {
	return &DefaultPrefix{
		Message: word,
	}
}
