package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
)

func getPage(url string) (int, error) {
	// see $godoc net/http Request | more
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	// see $godoc net/http Response | more
	defer resp.Body.Close()

	// see $godoc io/ioutil ReadAll | more
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func getter(url string, size chan string) {
	length, err := getPage(url)
	if err == nil {
		size <- fmt.Sprintf("%s has length %d", url, length)
	}
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
		"http://www.cnn.com",
	}

	size := make(chan string)

	for _, url := range urls {
		go getter(url, size)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-size)
	}
}
