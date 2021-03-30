package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func CountLines(pname string) (int, error) {
	f, err := os.Open(pname)
	if err != nil {
		return 0, err
	}
	r := bufio.NewReader(f)
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		n, err := r.Read(buf)
		count += bytes.Count(buf[:n], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

func main() {
	fmt.Println(CountLines("./core.go"))
}
