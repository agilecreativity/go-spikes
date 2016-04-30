// Package main provides ...
package main

import (
	"fmt"
	"net/http"
)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["name"][0]
	// // TODO: add the package
	// p, err := poetry.LoadPoem(poemName)
	// if err != nil {
	// 	http.Error(w, "File not found", http.StatusInternalServerError)
	// }
	fmt.Fprintf(w, "Poem is coming soon :%s\n", poemName)
}

func main() {
	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8088", nil)
}
