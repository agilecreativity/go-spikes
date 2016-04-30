package main

import (
	"fmt"
	"net/http"
)

const (
	ADDRESS = ":1024"
)

func main() {
	message := "Hello World"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
	})

	if e := http.ListenAndServe(ADDRESS, nil); e != nil {
		fmt.Println(e)
	}
}
