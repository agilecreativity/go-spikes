package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")
}

func main() {
	words := make([]string, 0, 4)
	fmt.Printf("%d %d\n", len(words), cap(words))
	words = append(words, "The")
	words = append(words, "quick")
	words = append(words, "brown")
	words = append(words, "fox")

	fmt.Printf("%d %d\n", len(words), cap(words))
	printer(words) // The quick brown fox

	// copy of the objects
	newWords := make([]string, 4)
	copy(newWords, words)
	newWords[2] = "blue"
	printer(newWords) // The quick blue fox

	// note the original words is unchanged
	printer(words) // The quick brow fox
}
