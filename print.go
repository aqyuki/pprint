package pprint

import "fmt"

const (
	// Space is char that space
	Space = 0x20
	// Error is prefix that error message
	Error = "Error"
)

// ErrorPrint print message with error prefix
func ErrorPrint(message string) {
	prefix := New(Error)
	printWithPrefix(message, prefix)
}

func printWithPrefix(message string, prefix Prefix) {
	fmt.Printf("%s%c%s\n", prefix.Get(), Space, message)
}
