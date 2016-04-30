package main

import (
	"fmt"
	"math"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, heigh, depth float64
	color               Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.heigh * b.depth
}

func (b Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BigestsColor() Color {
	v := 0.0
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

type Rectangle struct {
	width, heigh float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.heigh
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Are of r1 is :", r1.area())
	fmt.Println("Are of r2 is :", r2.area())
	fmt.Println("Are of c1 is :", c1.area())
	fmt.Println("Are of c2 is :", c2.area())
}
