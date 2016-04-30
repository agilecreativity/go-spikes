package main

import "fmt"

// see: https://gobyexample.com/for
func main() {
	// first single condition
	fmt.Println("simple for loop")
	i := 1
	for i <= 3 {
		i++
		fmt.Println(i)
	}

	fmt.Println("classic for loop")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("forever loop, stop by 'break'")
	for {
		fmt.Println("loop")
		break
	}
}
