package main

// from: effective-go (page 11)
import "fmt"
import "os"
import "io"

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
	var (
		s   string
		err error
	)
	s, err = Contents("read-file.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content: %s", s)
}
