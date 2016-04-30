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

var servers sync.WaitGroup

func init() {
	go SignalHander(make(chan os.Signal, 1))
}

func Launch(name string, f func() error) {
	servers.Add(1)
	go func() {
		defer servers.Done()
		if e := f(); e != nil {
			fmt.Println(name, "->", e)
			syscall.Kill(syscall.Getpid(), syscall.SIGABRT)
		}
	}()
}

func SignalHander(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)

	for s := <-c; ; s = <-c {
		switch s {
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

	Launch("HTTP", func() error {
		return http.ListenAndServe(ADDRESS, nil)
	})

	Launch("HTTPS", func() error {
		return http.ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})

	servers.Wait()
}
