package main

import "flag"
import "fmt"

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// Sample Usage:
// $go build 60-command-line-flags.go
// $./60-command-line-flags.go -word=opt -numb=7 -fork -svar=flag
// $./60-command-line-flags -word=opt
// $./60-command-line-flags -word=opt a1 a2 a3
