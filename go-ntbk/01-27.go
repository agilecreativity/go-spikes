package main

import "fmt"

const (
	Hello = "hello"
)

var world string

func init() {
	world = "world"
	fmt.Println(Hello, world)
}

func main() {
	// empty but should still produce "Hello world"
}
