package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, world")
	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output")
	case n != 12:
		fmt.Printf("No of bytes is not expect 12, but found %d", n)
	default:
		fmt.Printf("OK")
	}

	fmt.Printf("\n")
}
