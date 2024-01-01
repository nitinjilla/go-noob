package main

import (
	"errors"
	"fmt"
)

func main() {

	message, err := printer("Hello")

	if err != nil {
		fmt.Println(message)
	} else {
		fmt.Println(message)
	}
}

func printer(m string) (ms string, err error) {
	if m == "" {
		err = errors.New("Empty string! Function must contain a value of type string.")
	}

	ms = m
	return
}
