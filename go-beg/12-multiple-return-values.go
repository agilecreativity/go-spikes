// see: https://gobyexample.com/multiple-return-values
package main

import "fmt"

func plusAndMult(i, j int) (int, int) {
    return i + j, i * j
}

func values() (int, int) {
	return 3, 7
}

func main() {
	a, b := values()
	fmt.Println(a) //=> 3
	fmt.Println(b) //=> 7

	// ignore the first result
	_, c := values()
	fmt.Println(c) //=> 7

    x, y := plusAndMult(2, 3)
    fmt.Println(x, y) // => 5 6
}
