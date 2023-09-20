package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	//exported function
	if name == "" {
		return "", errors.New(" no name received ")
	}
	greet := fmt.Sprintf("Hello %v , Welcome! ", name)
	return greet, nil
}
