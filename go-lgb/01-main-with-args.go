package main

import (
    "fmt"
    "os"
)

func main() {
    // name, power := "Goku", 9000
    // fmt.Printf("%s's power is over %d\n", name, power)
    if len(os.Args) != 2 {
        os.Exit(1)
    }
    fmt.Println("It's over ", os.Args[1])
}

// $go run main-with-args.go 9000
