package main

import (
	"fmt"
)

func emit(c chan string) {
	words := []string{"The", "quick", "brown", "fox", "jump"}
	for _, word := range words {
		// send down c to the channel
		c <- word
	}
	// close
	close(c)
}

func main() {
	wordChannel := make(chan string)

	// producer of resources
	go emit(wordChannel)

	// consumer of resource
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")
}
