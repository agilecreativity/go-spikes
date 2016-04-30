// from http://go-talks.appspot.com/github.com/davecheney/presentations/introduction-to-go.slide#21

package main

import "fmt"

type Point struct {
    x, y int
}

type Weekday int

func (p Point) String() string {
  return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (d Weekday) String() string {
  return fmt.Sprintf("(%d)", d)
}

func main() {
  p := Point{2, 3}
  fmt.Println(p.String())
  fmt.Println(Point{4,5}.String())
  fmt.Println(Weekday{2 })
}

