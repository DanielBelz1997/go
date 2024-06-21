package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello  returns a greeting for the named person

func Hello(name string) (string, error) {
	if name == "" {
		return  name, errors.New("empty name")
	}
	// Return a greeting that embeds the name in the message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// did not touch the first function to nit harm anyone by changing the function
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hey, %v. Welcome!",
		"Great To See You Again, %v!",
		"Hail, %v! Well Met",
	}

	return formats[rand.Intn(len(formats))]
}
