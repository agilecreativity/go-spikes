package main

import (
	"fmt"
)

func emit(wordChan chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox"}

	i := 0
	for {
		select {
		case wordChan <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}

		case <-done:
			done <- true
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh)
	}

	// try stop after 100 iteration
	doneCh <- true
	<-doneCh
}
