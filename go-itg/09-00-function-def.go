package main

import "os"

func printer(msg string) (e error) {
	f, err := os.Create("helloworld.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	f.Write([]byte(msg))

	return err
}

func main() {
	printer("Hello, World!")
}
