package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	go func() {
		ch = make(chan int, 1)
		ch <- 1
		fmt.Println("w")
	}()
	go func(ch *chan int) {
		time.Sleep(time.Second)
		<-(*ch)
		fmt.Println("r")
	}(&ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
