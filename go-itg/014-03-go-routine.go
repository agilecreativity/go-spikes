package main

import (
	"fmt"
	"math/rand"
)

func makeRandoms(c chan int) {
	for {
		c <- rand.Intn(20)
	}
}

func main() {
	randoms := make(chan int)

	// producer of resources
	go makeRandoms(randoms)

	// receive from word channel
	for number := range randoms {
		fmt.Printf("%d ", number)
	}
}
