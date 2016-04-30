package main

import (
	"fmt"
)

func makeID(idChan chan int) {
	var id int
	id = 0
	for {
		idChan <- id
		id += 1
	}
}

func main() {
	idChan := make(chan int)

	go makeID(idChan)

	fmt.Printf("%d\n", <-idChan) // 0
	fmt.Printf("%d\n", <-idChan) // 1
	fmt.Printf("%d\n", <-idChan) // 2
	fmt.Printf("%d\n", <-idChan) // 3
}
