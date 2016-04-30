// see: https://gobyexample.com/arrays
package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty array:", a) // empty array [0 0 0 0 0]

	a[4] = 100
	fmt.Println("set:", a)    // set: [0 0 0 0100]
	fmt.Println("set:", a[4]) // set: 100

	fmt.Println("len:", len(a)) // len: 5

	b := [5]int{1, 2, 3, 4}
	fmt.Println("dcl:", b) // dcl: [1, 2, 3, 4, 0] // note: 0 is assigned to the last item

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", c) // dcl: [1, 2, 3, 4, 5]

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD) // dcl: [[0, 1, 2], [1, 2, 3]]
}
