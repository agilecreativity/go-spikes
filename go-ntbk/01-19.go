package main

import "fmt"

type HelloWorld struct{}

func (h HelloWorld) String() string {
	return "Hello World"
}

type Message struct {
	HelloWorld
}

func main() {
	m := &Message{}
	fmt.Println(m.HelloWorld.String())
	fmt.Println(m.String())
	fmt.Println(m)
}
