package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

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

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("The name is empty")
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
