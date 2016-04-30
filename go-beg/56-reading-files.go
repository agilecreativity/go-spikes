package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// this is the usual way to read the file before Go
	dat, err := ioutil.ReadFile("README.md")
	check(err)
	fmt.Print(string(dat))

	// However in Go it is nicer to do it this way
	f, err := os.Open("/tmp/dat")
	check(err)
	fmt.Println(f)

	// read the first 5 bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	b4, err := f.Seek(0, 0)
	check(err)
	fmt.Printf("...%s\n", string(b4))

	r5 := bufio.NewReader(f)
	b5, err := r5.Peek(5)

	check(err)
	fmt.Printf("5 bytes: %s\n", string(b5))

	f.Close()
}
