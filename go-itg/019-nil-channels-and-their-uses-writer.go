package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-t.C:
			ch = nil
		}
	}
}

func writer(ch chan int) {
	t := time.NewTimer(2 * time.Second)
	for {
		select {
		case ch <- rand.Intn(42):
		case <-t.C:
			ch = nil
		}
	}
}
func main() {
	intCh := make(chan int)
	go reader(intCh)
	go writer(intCh)
	time.Sleep(10 * time.Second)
}
