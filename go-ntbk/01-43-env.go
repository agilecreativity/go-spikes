package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

const (
	SECURE_ADDRESS = ":1025"
)

var address string

var servers sync.WaitGroup

func init() {
	if address = os.Getenv("SERVE_HTTP"); address == "" {
		address = ":1024"
	}
}

func Launch(f func()) {
	servers.Add(1)
	go func() {
		defer servers.Done()
		f()
	}()
}

func main() {
	message := "hello world"

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
	})

	Launch(func() {
		http.ListenAndServe(address, nil)
	})

	Launch(func() {
		http.ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})

	servers.Wait()
}
