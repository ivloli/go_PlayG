package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 3}
	fmt.Println(a[0:0])
	fmt.Println(a[3:])
}
