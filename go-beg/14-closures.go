// see: https://gobyexample.com/closures
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt()) //=> 1
	fmt.Println(nextInt()) //=> 2
	fmt.Println(nextInt()) //=> 3

	newInts := intSeq()
	fmt.Println(newInts()) //=> 1
	fmt.Println(newInts()) //=> 2
}
