package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var (
	address        string
	secure_address string
	certificate    string
	key            string
)

var servers sync.WaitGroup

func init() {
	if address = os.Getenv("SERVE_HTTP"); address == "" {
		address = ":1024"
	}

	if secure_address = os.Getenv("SERVE_HTTPS"); secure_address == "" {
		secure_address = ":1025"
	}

	if certificate = os.Getenv("SERVE_CERT"); certificate == "" {
		certificate = "cert.pem"
	}

	if certificate = os.Getenv("SERVE_KEY"); key == "" {
		key = "key.pem"
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
		http.ListenAndServeTLS(address, certificate, key, nil)
	})

	servers.Wait()
}
