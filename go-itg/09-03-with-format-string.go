package main

import "fmt"

func printer(format string, messages ...string) {
	for _, msg := range messages {
		fmt.Printf(format, msg)
	}
}
func main() {
	printer("%s\n", "Hello", "World", "And ..")
	// as hex
	printer("%x\n", "Hello", "World", "And ..")
}
