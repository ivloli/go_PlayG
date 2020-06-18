package main

import "fmt"

func go_test(c chan string) {
	<-c
	fmt.Println("AA")
}
func main() {
	c := make(chan string)
	for i := 0; i < 10; i++ {
		go go_test(c)
	}
	close(c)
}
