package main

import "fmt"

func main() {
	// multiple assignment in one line
	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
		fmt.Printf("Hello %d, and %d\n", i, j)
	}
	fmt.Println("-----------------")

	// classic for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - Hello, World!\n", i)
	}

	fmt.Println("-----------------")
	// simple for loop
	var counter int
	for counter < 10 {
		fmt.Printf("%d - Hello, World!\n", counter)
		counter += 1
	}

	fmt.Println("------------------")
	var stop bool

	i := 0

	for !stop {
		fmt.Printf("Forever loop : %d \n", i)
		i += 1
		if i == 10 {
			stop = true
		}
	}
}
