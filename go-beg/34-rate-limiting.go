package main

import "time"
import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

//=>
// request 1 2014-11-09 19:56:13.94556027 +1100 AEDT
// request 2 2014-11-09 19:56:14.142043106 +1100 AEDT
// request 3 2014-11-09 19:56:14.342864783 +1100 AEDT
// request 4 2014-11-09 19:56:14.543507295 +1100 AEDT
// request 5 2014-11-09 19:56:14.744139153 +1100 AEDT
// request 1 2014-11-09 19:56:14.744223675 +1100 AEDT
// request 2 2014-11-09 19:56:14.744248728 +1100 AEDT
// request 3 2014-11-09 19:56:14.744268407 +1100 AEDT
// request 4 2014-11-09 19:56:14.944789562 +1100 AEDT
// request 5 2014-11-09 19:56:15.145432882 +1100 AEDT
