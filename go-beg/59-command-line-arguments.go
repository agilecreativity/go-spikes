package main

import "os"
import "fmt"

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	args := os.Args[3]

	fmt.Println("1st:", argsWithProg)
	fmt.Println("2nd:", argsWithoutProg)
	fmt.Println("3rd:", args)
}

// To execute:
// $go build 59-command-line-arguments.go
// $./59-command-line-arguments a b c d
