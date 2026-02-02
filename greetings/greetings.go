package greetings

import (
	"fmt"
	"math/rand"

	"errors"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty Name")
	}


	message := fmt.Sprintf(randomFormet(),name)
	return  message, nil
} 


func randomFormet() string {
	formats := []string{
		"hi %v bro",
		"Hello %v",
	}


	return formats[rand.Intn(len(formats))]
}