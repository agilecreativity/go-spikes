package main

import "fmt"

type Stringer interface {
	String() string
}

type Hello struct{}

func (h Hello) String() string {
	return "Hello"
}

type World struct{}

func (w *World) String() string {
	return "world"
}

type Message struct {
	X fmt.Stringer
	Y fmt.Stringer
}

func (v Message) IsGreeting() (ok bool) {
	if _, ok = v.X.(*Hello); !ok {
		_, ok = v.Y.(*Hello)
	}
	return
}

func main() {
	m := &Message{}
	fmt.Println("empty interface:", m.IsGreeting())

	m.X = new(Hello)
	fmt.Println("add new(Hello):", m.IsGreeting())

	m.Y = new(World)
	fmt.Println("add new(World):", m.IsGreeting())

	m.Y = m.X
	fmt.Println("add m.Y = m.X:", m.IsGreeting())

	m = &Message{X: new(World), Y: new(Hello)}
	fmt.Println("swap X, and Y:", m.IsGreeting())

	m.X, m.Y = m.Y, m.X
	fmt.Println("swap :", m.IsGreeting())
}
