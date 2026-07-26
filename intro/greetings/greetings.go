// Package greetings contains a bunch a functions for greeting people
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello greets person by given name
// Fails if the name is empty
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos greats randomly multiple persons by names
// Fails if a name is empty
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
		"Hi, %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
