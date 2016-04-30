package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	if connection, e := net.Dial("tcp", ":1024"); e == nil {
		defer connection.Close()
		if text, e := bufio.NewReader(connection).ReadString('\n'); e == nil {
			fmt.Println("Hello")
			fmt.Printf(text)
		}
	} else {
		fmt.Println("Something is not right!")
	}
}
