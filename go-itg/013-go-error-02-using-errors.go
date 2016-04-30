package main

import (
	"errors"
	"fmt"
	"os"
)

// use the errors from the errors package
var (
	errorEmptyString = errors.New("Unwilling to print an empty string")
)

func printer(msg string) error {
	if msg == "" {
		// Or if it is critical you can
		// panic(errorEmptyString)
		return errorEmptyString
	}

	_, err := fmt.Printf("%s\n", msg)

	return err
}

func main() {
	// Make it generate error condition intentionally
	if err := printer(""); err != nil {

		if err == errorEmptyString {
			fmt.Printf("You tried to print an empty string!")
		} else {
			fmt.Printf("Print failed: %s\n", err)
		}

		os.Exit(1)
	}
}
