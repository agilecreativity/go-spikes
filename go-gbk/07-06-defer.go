package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()

	// use case
	//f, _ := os.Open(filename)
	//defer f.Close()
}

// Result:
// 1st
// 2nd
