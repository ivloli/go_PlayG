package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func twoSum(a []int, t int) (i, j int) {
	if len(a) == 0 {
		return -1, -1
	}
	m := make(map[int]int, len(a))
	for index, value := range a {
		if v, ok := m[t-value]; ok {
			return index, v
		} else {
			m[value] = index
		}
	}
	return -1, -1
}
func twoDiff(a []int, t int) (i, j int) {
	if len(a) == 0 {
		return -1, -1
	}
	m := make(map[int]int, len(a))
	for index, value := range a {
		if v, ok := m[value-t]; ok {
			return index, v
		} else if v, ok := m[value+t]; ok {
			return index, v
		} else {
			m[value] = index
		}
	}
	return -1, -1
}
