package main

import "fmt"

func printer(messages ...string) {
	for _, msg := range messages {
		fmt.Printf("%s\n", msg)
	}
}
func main() {
	printer("Hello", "World", "And ..")
}
