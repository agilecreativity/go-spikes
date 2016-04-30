package main

import (
  "fmt"
  "time"
)

func printer(msg string, stopCh chan bool) {
    for {
        select {
        case <-stopCh:
            return
        default:
           fmt.Printf("%s\n", msg)
        }
    }
}

func main() {
    stopCh := make(chan bool)

    for i:= 0; i < 10; i++ {
        go printer(fmt.Sprintf("printer: %d", i), stopCh)
    }

    time.Sleep(2 * time.Second)
    close(stopCh)
    fmt.Println("FYI: after the channel close!")
    time.Sleep(2 * time.Second)
}
