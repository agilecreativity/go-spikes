package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

//=>
// Tick at 2014-11-09 19:43:31.873911365 +1100 AEDT
// Tick at 2014-11-09 19:43:32.373182138 +1100 AEDT
// Tick at 2014-11-09 19:43:32.873863802 +1100 AEDT
// Ticker stopped
