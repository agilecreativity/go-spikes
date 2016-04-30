package main

import "fmt"

// note: this is the slice
func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")
	// Note: this will be effect as we are using slice
	w[2] = "blue"
}

func main() {
	words := []string{"the", "quick", "brown", "fox"}
	fmt.Printf("%d\n", len(words))
	printer(words) //=> the quick brown fox
	// Now the result
	printer(words) //=> the quick blue fox

	printer(words[0:3]) //=> the quick blue fox
}
