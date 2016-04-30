package main

// see: http://www.slideshare.net/mstine/java-devlearnstogooscon
import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println("from", ":", i)
	}
}

func main() {
	// synchronous
	f("direct")

	go f("goroutine") // asynchronous

	go func(msg string) { // asynchronous and anonymous
		fmt.Println("Your message in go routine:", msg)
	}("hello")

	// Stop and see or else will not be able to see
	var input string
	fmt.Scanln(&input)
}
