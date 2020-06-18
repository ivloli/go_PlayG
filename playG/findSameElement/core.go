package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{5, 6, 7, 8, 9, 9, 9, 10, 22, 22, 22, 22, 31, 32, 33, 34, 34, 34, 44, 45, 45, 56, 56, 78, 78, 79, 81, 89, 100}
	b := []int{1, 2, 3, 9, 23, 45, 120}
	fmt.Println(bin(a, b))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func simple(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}
	res := make([]int, 0, min(len(a), len(b)))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func rightBound(a []int, s, e, t int) int {
	if e <= s {
		return -1
	}
	left, right := s, e
	mid := 0
	for left < right {
		mid = left + (right-left)/2
		if a[mid] == t {
			left = mid + 1
		} else if a[mid] < t {
			left = mid + 1
		} else if a[mid] > t {
			right = mid
		}
	}
	return right - 1
}

func bin(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}
	res := make([]int, 0, min(len(a), len(b)))
	s := 0
	e := len(a)
	for _, item := range b {
		if item <= a[s] {
			if item == a[s] {
				res = append(res, item)
				s++
			}
			if s == e {
				return res
			}
			continue
		}
		fmt.Println(s, item)
		r := rightBound(a, s, e, item)
		if a[r] == item {
			res = append(res, item)
		}
		s = r + 1
		if s >= e {
			return res
		}
	}
	return res
}
