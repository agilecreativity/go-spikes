package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	ADDRESS        = ":1024"
	SECURE_ADDRESS = ":1025"
)

var (
	address        string
	secure_address string
	certificate    string
	key            string
)

var servers sync.WaitGroup

func init() {
	go SignalHander(make(chan os.Signal, 1))
}

func Launch(f func()) {
	servers.Add(1)
	go func() {
		defer servers.Done()
		f()
	}()
}

func SignalHander(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)

	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("interrupt - continue running")
		case syscall.SIGABRT:
			fmt.Println("abnormal exit")
			os.Exit(1)
		case syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("clean shutdown")
			os.Exit(0)
		}
	}
}

func main() {
	message := "hello world"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
	})

	Launch(func() {
		http.ListenAndServe(ADDRESS, nil)
	})

	Launch(func() {
		http.ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})

	servers.Wait()
}
