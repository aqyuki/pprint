package pprint

type (
	Prefix interface {
		// Get return registered prefix
		Get() string
		// RegisterWord register new prefix
		RegisterWord(words ...string)
		// SetFormat set prefix format
		// Args :
		//  - format string
		// 		- give format string
		SetFormat(format string)
	}
)
