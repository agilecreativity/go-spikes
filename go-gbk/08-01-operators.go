package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPtr(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	zeroPtr(&x)
	fmt.Println(x) // x is now 0

	// using new
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // 1
}
