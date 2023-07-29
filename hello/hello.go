package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message and print it.
	message, err := greetings.Hello("hg")
	// If an error was returned, print it to the consle
	// and exist the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the consle
	fmt.Println(message)
}
