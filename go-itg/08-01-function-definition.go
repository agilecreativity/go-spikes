package main

import "fmt"

func printer(msg string) {
	fmt.Printf("%s\n", msg)
}

// function that returns error
func message_with_error(msg string) error {
	_, e := fmt.Printf("Your message: %s\n", msg)
	return e
}

// multiple return without the name:w
func message_multiple_result(msg string) (string, error) {
	msg += "\n"

	_, err := fmt.Printf(msg)
	return msg, err
}

// message with defer statement
func message_with_defer(msg string) error {

	defer fmt.Printf("First  -->1\n")
	defer fmt.Printf("Second -->2\n")

	_, err := fmt.Printf("%s\n", msg)

	return err
}

func main() {
	printer("Hello World")

	err := message_with_error("Hello")

	if err != nil {
		fmt.Printf("Found errors", err)
		panic(err)
	}

	msg, err := message_multiple_result("Good Bye")

	if err == nil {
		fmt.Printf("Another message result: %s\n", msg)
	}

	fmt.Printf("-------------->>>")
	message_with_defer("Message with defer")
}
