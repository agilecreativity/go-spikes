package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println(i) //=> 1

	zeroval(i)
	fmt.Println(i) //=> 1

	zeroptr(&i)
	fmt.Println(i) //=> 0

	fmt.Println(&i) //=> 0x21342423
}
