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
	X Stringer
	Y Stringer
}

func (v Message) String() (r string) {
	switch {
	case v.X == nil && v.Y == nil:
	case v.X == nil:
		r = v.Y.String()
	case v.Y == nil:
		r = v.X.String()
	default:
		r = fmt.Sprintf("%v %v", v.X, v.Y)
	}
	return
}

func main() {
	m := &Message{}
	fmt.Println("empty interface:", m)

	m.X = new(Hello)
	fmt.Println("add new(Hello):", m)

	m.Y = new(World)
	fmt.Println("add new(World):", m)

	m.Y = m.X
	fmt.Println("add m.Y = m.X:", m)

	m = &Message{X: new(World), Y: new(Hello)}
	fmt.Println("swap X, and Y:", m)

	m.X, m.Y = m.Y, m.X
	fmt.Println("swap :", m)
}
