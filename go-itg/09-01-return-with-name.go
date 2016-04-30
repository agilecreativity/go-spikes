// Package main
package main

import "fmt"

func printer(msg string) (e error) {
	_, e = fmt.Printf("%s\n", msg)
	// note: no need to use return e
	return
}

func printer_multiple_returns(msg string) (returnMsg string, e error) {
	// implicitly returned the 2nd parameter
	_, e = fmt.Printf("%s\n", msg)

	// Note: explicite return the first parameter
	returnMsg = "I did it"
	return
}

func main() {
	printer("Hello, World..")
	msg, err := printer_multiple_returns("Hello, World..")
	if err == nil {
		fmt.Printf("Your message %s\n", msg)
	}
}
