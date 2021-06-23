package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	mapt()
	fmt.Println("func fin1")
	time.Sleep(time.Second * 3)
	mapt()
	fmt.Println("func fin2")
	mapt()
	time.Sleep(time.Second * 3)
	fmt.Println("func fin3")
	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
}

func mapt() {
	mmap := make(map[int]string)
	for i := 0; i <= 1000000; i++ {
		mmap[i] = strconv.Itoa(i) + "this is a tail string"
	}
	mmap = nil
}
