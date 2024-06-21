package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Daniel", "Jeffry", "Assam"}


	message, err := greetings.Hellos(names)
	if (err != nil) {
		log.Fatal(err)
	}

	fmt.Println(message)
}



