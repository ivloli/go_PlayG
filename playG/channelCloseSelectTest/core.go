package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("vim-go")
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	outCh := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			outCh <- i
		}
		close(outCh)
	}()
	for {
		select {
		case <-interrupt:
			fmt.Println("exit")
			return
		case v, ok := <-outCh:
			if ok {
				fmt.Println(v)
				time.Sleep(time.Second)
			} else {
				return
			}
		}
	}
}
