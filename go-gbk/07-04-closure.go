package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1)) // 2
	fmt.Println(add(1, 3)) // 4

	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3

	//--------------------------
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}
