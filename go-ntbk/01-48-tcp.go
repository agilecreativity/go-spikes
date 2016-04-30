package main

import (
	"fmt"
	"net"
)

func main() {
	if listener, e := net.Listen("tcp", ":1024"); e == nil {
		for {
			if connection, e := listener.Accept(); e == nil {
				go func(c net.Conn) {
					defer c.Close()
					// TODO: may be use it to show the current time!
					fmt.Fprintf(c, "Hello, world")
				}(connection)
			}
		}
	}
}
