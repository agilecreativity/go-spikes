package main

import "fmt"

type empty interface{}
type example interface {
	notImplemented()
}

func main() {
	one := 1
	var i empty = one
	var float float32
	float = float32(one)

	switch i.(type) {
	case int:
		fmt.Printf("%d\n", i)
	default:
		fmt.Println("Type error\n")
	}

	fmt.Printf("%f\n", float)
}
