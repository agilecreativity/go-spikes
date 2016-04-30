package main
import (
    "bytes"
    "log"
    "os/exec"
)

func main() {
    path, err := exec.LookPath("ruby")

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Ruby can be found at: %s", path)

    // cmd := exec.Command("ruby", "-v")
    cmd := exec.Command(path, "-v")
    var out bytes.Buffer
    cmd.Stdout = &out

    err = cmd.Run()
    if err != nil {
        log.Fatal(err)
    }

    log.Print(out.String())
}
