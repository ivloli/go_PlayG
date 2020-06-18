package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	a := []int{1, 7, 6, 5, 4, 3, 2}
	b := []int{7, 1, 2, 3, 4, 5, 6}
	fmt.Println(core_next(a))
	fmt.Println(core_prev(b))
}
func reverse(a []int) {
	if len(a) < 2 {
		return
	}
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}
func core_next(a []int) []int {
	if len(a) < 2 {
		return a
	}
	i := len(a) - 1
	for i > 0 {
		if a[i-1] < a[i] {
			break
		}
		i--
	}
	if i == 0 {
		reverse(a)
		return a
	} else {
		i--
		j := i + 1
		for j < len(a) {
			if a[j] < a[i] {
				break
			}
			j++
		}
		a[i], a[j-1] = a[j-1], a[i]
		reverse(a[i+1:])
		return a
	}
}
func core_prev(a []int) []int {
	if len(a) < 2 {
		return a
	}
	i := len(a) - 1
	for i > 0 {
		if a[i-1] > a[i] {
			break
		}
		i--
	}
	if i == 0 {
		reverse(a)
		return a
	}
	i--
	t := len(a) - 1
	for t > i {
		if a[t] < a[i] {
			break
		}
		t--
	}
	a[t], a[i] = a[i], a[t]
	reverse(a[i+1:])
	return a
}
