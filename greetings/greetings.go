package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// := operator is a shortcut for declaring and initializing a variable in one line.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// nil (meaning no error) as a second value in the successful return.
	return message, nil
}
