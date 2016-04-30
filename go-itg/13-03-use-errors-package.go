package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errorEmptyString = errors.New("Unwilling to print an empty string")
)

func printer(msg string) error {
	if msg == "" {
		// custom error
		// return errorEmptyString
		panic(errorEmptyString)
	}

	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer(""); err != nil {
		if err == errorEmptyString {
			fmt.Printf("You try to print an empty string")
		} else {
			fmt.Printf("Other error failed: %s\n", err)
		}

		os.Exit(1)
	}
}
