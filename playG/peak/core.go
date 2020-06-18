package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func core(a []int) int {
	l := len(a)
	if l == 0 {
		return -1
	}
	if l == 1 {
		return 0
	}
	if a[0] > a[1] {
		return 0
	}
	if a[l-1] > a[l-2] {
		return l - 1
	}
	i, j := 0, l-1
	mid := 0
	for i < j {
		mid = i + (j-i)/2
		if a[mid] < a[mid+1] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}
