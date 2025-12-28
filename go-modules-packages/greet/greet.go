package greet

import "fmt"

// Hello returns a greeting message for the given name.
func Hello(name string) string {
	message := fmt.Sprintf("Hello, %s! Welcome to Go packages.", name)
	return message
}

// Goodbye returns a farewell message for the given name.
func Goodbye(name string) string {
	message := fmt.Sprintf("Goodbye, %s! See you again.", name)
	return message
}
