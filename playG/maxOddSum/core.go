package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(core(10))
}

func core(a int64) int64 {
	var res int64
	for i := a; i > 0; i = i / 2 {
		var temp int64
		temp = (i + 1) / 2
		res += temp * temp
	}
	return res
}
