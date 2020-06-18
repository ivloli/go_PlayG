package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	fmt.Println("vim-go")
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	for i := 0; i < 10; i += 2 {
		go func(i int) {
			defer fmt.Printf("defer: %d \n", i)
			i++
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
