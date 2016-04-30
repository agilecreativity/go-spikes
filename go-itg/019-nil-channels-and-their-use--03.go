package main

import (
 "fmt"
 "time"
 "math/rand"
)

func reader(ch chan int) {
    t := time.NewTimer(3 * time.Second)
    for {
        select {
        case i := <- ch:
            fmt.Printf("%d\n", i)
        case <- t.C:
            ch = nil
        }
    }
}

func writer(ch chan int) {

    stopper   := time.NewTimer(2 * time.Second)
    resterter := time.NewTimer(5 * time.Second)

    savedCh := ch

    for {
       select {
       case ch <- rand.Intn(42):
       case <- stopper.C:
           ch = nil
       case <- resterter.C:
           ch = savedCh
       }
    }
}
func main() {
  intCh := make(chan int)
  go reader(intCh)
  go writer(intCh)

  // make sure it runs long enough
  time.Sleep(30 * time.Second)
}
