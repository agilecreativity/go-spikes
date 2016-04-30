package main

import "fmt"

func printer(w [4]string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	// won't take effect as array is passed by value
	w[2] = "blue"
}

func main() {
	words := [...]string{"the", "quick", "brown", "fox"}
	fmt.Printf("%s\n", words[2]) // 'brown'

	words = [4]string{"the", "quick", "brown", "fox"}
	printer(words)

	printer(words)
}
