package pprint

import (
	"fmt"
)

const (
	// Space is char that space
	Space = 0x20
	// Error is prefix that error message
	Error = "Error"
	// Info is prefix that info message
	Info = "Info"
)

// ErrorPrint print message with error prefix
func ErrorPrint(message string) {
	prefix := New(Error)
	printWithPrefix(message, prefix)
}

// InfoPrint print message with information prefix
func InfoPrint(message string) {
	prefix := New("Info")
	printWithPrefix(message, prefix)
}

func printWithPrefix(message string, prefix Prefix) {
	fmt.Printf("%s%c%s\n", prefix.Get(), Space, message)
}
