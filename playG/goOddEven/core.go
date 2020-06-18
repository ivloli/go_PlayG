package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	N := 100
	c1 := make(chan bool, 1)
	c2 := make(chan bool)
	done := make(chan bool)
	go func(N int) {
		for i := 1; i <= N; i += 2 {
			<-c1
			fmt.Println("fist:", i)
			c2 <- true
		}
		close(c2)
	}(N)
	go func(N int) {
		for i := 2; i <= N; i += 2 {
			<-c2
			fmt.Println("second:", i)
			c1 <- true
		}
		close(c1)
		close(done)
	}(N)
	c1 <- true
	<-done
}
