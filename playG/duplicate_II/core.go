package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 4, 5, 6}
	r := removeDuplicates(a)
	fmt.Println(a)
	fmt.Println(r)
	b := []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 4, 5, 6}
	fmt.Println(removeDuplicates_II(b))
	fmt.Println(b)
}
func removeDuplicates(a []int) int {
	if len(a) < 2 {
		return len(a)
	}
	i := 0
	for _, item := range a[1:] {
		if a[i] != item {
			i++
			a[i] = item
		}
	}
	return i + 1
}

func removeDuplicates_II(a []int) int {
	if len(a) < 3 {
		return len(a)
	}
	i := 0
	for _, item := range a {
		if i < 2 || item != a[i-2] {
			a[i] = item
			i++
		}
	}
	return i
}
