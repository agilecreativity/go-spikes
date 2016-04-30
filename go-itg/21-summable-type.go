package main

import "fmt"

type SummableSlice []int

func (s SummableSlice) sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {
	var s SummableSlice = SummableSlice{1, 2, 3, 4, 5}
	fmt.Printf("Sum of %v is %d\n", s, s.sum())
}
