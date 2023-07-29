package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// := operator is a shortcut for declaring and initializing a variable in one line.
	message := fmt.Sprintf(randomFormat(), name)
	// nil (meaning no error) as a second value in the successful return.
	return message, nil
}

// randomFormat starts with a lowercase letter, making it accessible only to code
// in its own package (in other words, it's not exported).
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
