package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{0, 0, 0, 0, 0, 0}
	fmt.Println(a)
	fmt.Println(core(a))
	b := []int{0, 0, 0, 1, 2, 3, 0, 0, 4, 0, 5, 5, 0}
	fmt.Println(b)
	fmt.Println(core(b))
	c := []int{1, 2, 3, 4, 5, 6, 7, 0, 0, 7, 0}
	fmt.Println(c)
	fmt.Println(core(c))
}

func core(a []int) []int {
	if len(a) == 0 {
		return a
	}
	i := -1
	for j := 0; j < len(a); j++ {
		if a[j] != 0 {
			i++
			if i != j {
				a[i] = a[j]
				a[j] = 0
			}
		}
	}
	return a
}
