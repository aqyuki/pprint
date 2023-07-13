package pprint

import (
	"fmt"

	"github.com/aqyuki/pprint/custom"
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

// CustomPrint print message with customized prefix
func CustomPrint(message string, prefix custom.CustomPrefix) {
	printWithPrefix(message, prefix)
}

func printWithPrefix(message string, prefix Prefix) {
	fmt.Printf("%s%c%s\n", prefix.Get(), Space, message)
}
