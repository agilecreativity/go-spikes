package main

import (
	"fmt"
	"net/http"
)

const (
	SECURE_ADDRESS = ":1024"
)

func main() {
	message := "Hello World"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
	})

	http.ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
}

// Note: to generate the key before trying this
// $go run $GOROOT/src/pkg/crypto/tls/generate_cert.go -ca=true -host="localhost"
