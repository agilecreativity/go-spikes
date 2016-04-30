package main

import (
	"fmt"
	"time"
)

// send string through channel (sender)
// producer/sender
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// another channel
func ponger(c chan string) {
	// loop forever
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// consumer/receiver of the channel
func printer(c chan string) {
	for {
		// receive a message and store it in `msg`
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
