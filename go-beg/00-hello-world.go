package main

// import "fmt"
// import "os"
// import "io"

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer f.Close()

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func main() {
	fmt.Println("hello world")

	for pos, char := range "สวัสดี" {
		fmt.Printf("Char %c starts at byte %d\n", char, pos)
	}
	var s string
	s, _ = Contents("00-hello-world.go")
	fmt.Println("Content: %s", s)
}

//$go run 00-hello-world.go
//$go build 00-hello-world.go
//$./00-hello-world
