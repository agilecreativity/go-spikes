package main

import (
	"fmt"
)

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func some_func(f int) (int, int) {
	return f * 10, f + 10
}

// variable argument functions
// func Println(a ...interface{}) (n int, err error) { }

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	xs := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(average(xs)) // 5.5

	x, y := some_func(10)
	fmt.Println("x = ", x, "y = ", y)

	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// special way with slice
	ys := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(add(ys...))
}
