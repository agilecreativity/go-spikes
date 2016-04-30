package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

// type Printer func(string) //()
type Printer func(string)

func Greet(s Salutation, do Printer) {
	name, msg := CreateMessage(s.name, s.greeting, "Yo!")
	do(name)
	do(msg)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message, alternate = name+" "+greeting[0], "HEY "+name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func PrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}

func main() {
	var s = Salutation{"Bob", "Hi"}
	// Greet(s, PrintLine)
	// Greet(s, Print)
	Greet(s, CreatePrintFunction("!!!"))
}
