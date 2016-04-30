// see: http://talks.golang.org/2014/taste.slide#7
package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprintf(w, "Hello! บุรินทร์\nURL= %s\n", req.URL)
}

func main() {
	fmt.Println("Please connect to localhost:7777/hello")
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
