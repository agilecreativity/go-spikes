package main
import (
   "bytes"
   "log"
   "os/exec"
)

func main() {
    // cmd := exec.Command("ruby", "-e", `puts 42*42`)
    // cmd := exec.Command("ruby", "-e", `puts "What's your name?"; name = gets.chomp; puts "Hi #{name}!"`)
    // cmd.Stdin = strings.NewReader("Burin")

    cmd := exec.Command("ruby", "-e", `sleep(5); puts "done sleeping";`)

    var out bytes.Buffer
    cmd.Stdout = &out

    // err := cmd.Run()
    err := cmd.Start()

    if err != nil {
        log.Fatal(err)
    }
    log.Println("Waiting for command")

    err = cmd.Wait() // blocking call
    if err != nil {
        log.Fatal(err)
    }

    log.Print(out.String())
}
