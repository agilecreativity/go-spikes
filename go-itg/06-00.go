package main

import (
	"fmt"
	"os"
)

func main() {
	byteCount, err := fmt.Printf("Hello, World\n")

	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Number of bytes read %d\n", byteCount)
}
