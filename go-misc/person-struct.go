// see: https://talks.golang.org/2014/go4java.slide#46
package main

import "fmt"

type Person struct{ Name string }

func (p Person) Introduce() { fmt.Println("Hi, I am", p.Name) }

type Employee struct {
	Person
	EmployeeID int
}

func main() {
	var e Employee
	e.Name = "Peter"
	e.EmployeeID = 1234
	e.Introduce()
}
