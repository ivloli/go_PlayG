package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(core(3, 10, 5))
}

func core(box, fish int, max int) int {
	if box == 0 {
		return 0
	}
	if fish < 0 {
		return 0
	}
	if box == 1 {
		if fish > max {
			return 0
		} else {
			return 1
		}
	}
	if fish == 0 {
		return 1
	}

	var res int
	for i := 0; i <= max; i++ {
		res += core(box-1, fish-i, max)
	}
	return res
}
