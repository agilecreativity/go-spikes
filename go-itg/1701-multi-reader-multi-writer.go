package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
		"http://www.abc.com.au",
	}
	for _, url := range urls {
		pageLength, err := getPage(url)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%s is length %d\n", url, pageLength)
	}
}
