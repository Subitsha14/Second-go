package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	//exported function
	if name == "" {
		return "", errors.New(" no name received ")
	}
	greet := fmt.Sprintf(randomFormat(), name)
	return greet, nil
}

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
	formats := []string{"Hii %v. Wecome! \n", "Great to see you, %v \n", "Hail, %v! well met! \n"}

	return formats[rand.Intn(len(formats))]
}
