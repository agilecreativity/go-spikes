package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

func logger(pipe io.ReadCloser) {
	reader := bufio.NewReader(pipe)

	for {
		output, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
			return
		}

		log.Print(string(output))
	}
}

func main() {
	cmd := exec.Command("vim", "test.txt")

	pipe, err := cmd.StderrPipe()

	if err != nil {
		log.Fatal(err)
	}

	go logger(pipe)

	err = cmd.Run()

	log.Fatal(err)
}
