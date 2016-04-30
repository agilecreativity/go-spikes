package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer f.Close()

	// read from the file
	b := make([]byte, 100)
	n, err := f.Read(b)

	// byte slice
	fmt.Printf("%d: %x \n", n, b)

	// string version (Note this is the copy not a reference value)
	stringVersion := string(b)
	fmt.Printf("%d: %s \n", n, stringVersion)

	// the patterns that is quite common
	// someString := "foo bar"
	// f.Write([]byte(someString))
}
