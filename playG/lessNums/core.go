package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
	a := []int{8, 1, 2, 2, 3}
	fmt.Println(core(a))
}

func core(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b)
	m := make(map[int]int)
	for i, v := range b {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = i
	}
	for i, v := range a {
		b[i] = m[v]
	}
	return b
}
