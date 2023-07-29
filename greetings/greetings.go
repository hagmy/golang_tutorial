package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// := operator is a shortcut for declaring and initializing a variable in one line.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
