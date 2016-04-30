package main

import "fmt"
import "os"

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

// Usage:
// $go run 65-exit.go
//=> exit status 3

// $go build 65-exit.go
// $./65-exit
// $echo $?
//=> 3
