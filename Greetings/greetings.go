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

func randomFormat() string {
	formats := []string{"Hii %v. Wecome!", "Great to see you, %v", "Hail, %v! well met!"}

	return formats[rand.Intn(len(formats))]
}
