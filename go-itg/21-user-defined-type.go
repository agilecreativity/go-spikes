package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}

	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		w.err = err
	}
}

// function thats return value
func (w *webPage) isOK() bool {
	return w.err == nil
}

func main() {
	// w := &webPage{}
	w := new(webPage)

	w.url = "http://www.oreilly.com/"

	w.get()
	if w.isOK() {
		fmt.Printf("URL: %s, length: %d\n", w.url, len(w.body))
	} else {
		fmt.Printf("Something went wrong: %s\n", w.err)
	}
}
