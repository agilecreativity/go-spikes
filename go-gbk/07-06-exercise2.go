package main

import "fmt"

// half the integer and indicate if it is even
func half(i int) (int, bool) {
	return i / 2, i%2 == 0
}

func max(args ...int) int {
	max := 0
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	x, y := half(4)
	fmt.Println(x) // 2
	fmt.Println(y) // true

	a, b := half(5)
	fmt.Println(a) // 2
	fmt.Println(b) // true

	// using slice
	fmt.Println(max([]int{1, 2, 10, 4, 6, 9}...)) // 10

	// using normal way
	fmt.Println(max(1, 2, 10, 4, 6, 9)) // 10
}
