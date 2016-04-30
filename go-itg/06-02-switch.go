package main

import "fmt"

func main() {
	atoz := "the quick brown fox jumps over the lazy dog"
	vowels := 0
	consonants := 0
	zeds := 0
	for _, r := range atoz {

		switch r {
		// no fall through in Go
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		case 'z':
			zeds += 1
			fallthrough
		default:
			consonants += 1
		}
	}

	fmt.Printf("Vowels count :%d\n", vowels)
	fmt.Printf("Consonant count :%d\n", consonants)
	fmt.Printf("Z count :%d\n", zeds)
}
