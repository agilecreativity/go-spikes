package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// binary, lookErr := exec.LookPath("ls")
	binary, lookErr := exec.LookPath("vim")

	if lookErr != nil {
		panic(lookErr)
	}

	filename := "40-panic.go"

	args := []string{"-E", "-c 'TOhtml'", "-c 'w'", "-c 'qa!'", filename, "> /dev/null"}
	fmt.Println("args:", args)

	// args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	// fmt.Println(env)
	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}
}
