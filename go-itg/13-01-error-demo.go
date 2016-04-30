package main

import (
	"fmt"
	"os"
)

func printer(msg string) error {
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	// err := printer("some message")
	// if err != nil {
	//    os.Exit(1)
	// }

	if err := printer("Hello, World!"); err != nil {
		os.Exit(1)
	}
}
