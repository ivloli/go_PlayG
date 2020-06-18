package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{3, 3, 3, 3, 3, 3, 3}
	b := []int{3, 3, 3, 3, 4, 5, 6, 1, 2, 3, 7}
	fmt.Println(core(a))
	fmt.Println(core(b))
}

func core(a []int) int {
	if len(a) < 2 {
		return -1
	}
	max, smax := 0, -1
	for i, v := range a {
		if v > a[max] {
			smax = max
			max = i
		} else if v < a[max] && (smax == -1 || v > a[smax]) {
			smax = i
		}
	}
	return smax
}
