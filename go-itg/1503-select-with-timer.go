package main

import (
	"fmt"
	"time"
)

func emit(wordChan chan string, done chan bool) {
	defer close(wordChan)
	words := []string{"The", "quick", "brown", "fox"}

	i := 0
	t := time.NewTimer(1 * time.Second)

	for {
		select {
		case wordChan <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}

		case <-done:
			done <- true

		case <-t.C:
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
}
