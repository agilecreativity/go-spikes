package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Println("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true) // true
	fmt.Printf("%d\n", 123)  // 123
	fmt.Printf("%b\n", 14)   // 1110
	fmt.Printf("%c\n", 33)   // !
	fmt.Printf("%x\n", 456)  // 1c8
	fmt.Printf("%f\n", 78.9) // 78.900000
	fmt.Printf("%e\n", 1234000000.0)
	fmt.Printf("%E\n", 1234000000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Fprintf(os.Stderr, "an %s\n", "error..")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
}
