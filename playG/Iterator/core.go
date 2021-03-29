package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
	f, err := os.Open("test")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		reader := bufio.NewReader(f)
		for line, _, err := reader.ReadLine(); err == nil; {
			fmt.Println(string(line))
			line, _, err = reader.ReadLine()
		}
	*/
	/*
		it := Iterator(f)
		for {
			line, ok := it()
			if !ok {
				break
			}
			fmt.Println(line)
		}
	*/
	ch := IteratorWithChan(f)
	for line := range ch {
		fmt.Println(line)
	}
	return
}

func Iterator(f *os.File) (fn func() (string, bool)) {
	if f == nil {
		return nil
	}
	reader := bufio.NewReader(f)
	return func() (string, bool) {
		line, _, err := reader.ReadLine()
		if err != nil {
			return "", false
		} else {
			return string(line), true
		}
	}
}

func IteratorWithChan(f *os.File) <-chan string {
	if f == nil {
		return nil
	}
	reader := bufio.NewReader(f)
	ch := make(chan string)
	go func() {
		for line, _, err := reader.ReadLine(); err == nil; {
			ch <- string(line)
			line, _, err = reader.ReadLine()
		}
		close(ch)
		fmt.Println("end")
	}()
	return ch
}
