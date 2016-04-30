package main

import "fmt"

type HelloWorld bool

func (h HelloWorld) String() (r string) {
	if h {
		r = "Hello World"
	}
	return
}

type Message struct {
	HelloWorld
}

func main() {
	m := &Message{HelloWorld: true}
	fmt.Println("first one", m)
	m.HelloWorld = false
	fmt.Println("after set to false", m)
	m.HelloWorld = true
	fmt.Println("after set to true", m)
}
