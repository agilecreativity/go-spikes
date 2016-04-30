package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
}

// interface as fields:
type MultiShape struct {
	shapes []Shape
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// helper methods
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// one way to implement this
func circleArea1(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// better way to do this
func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// best to use the methods with name and type
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 5}

	fmt.Println(circleArea1(c)) // 78.53981633974483

	fmt.Println(circleArea2(&c)) // 78.53981633974483

	// now we can call the method like so, no longer need the & operator
	fmt.Println(c.area()) // 78.53981633974483

	// now call one for rectangle
	r := Rectangle{0, 0, 10, 20}
	fmt.Println(r.area()) // 200

	fmt.Println(totalArea(&c, &r)) // 278.53981633974483

	// and use it as this way
	shapes := MultiShape{[]Shape{&c, &r}}
	fmt.Println(shapes.area())
}
