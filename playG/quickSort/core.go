package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := make([]int, 100)
	rand.Seed(time.Now().UnixNano())
	for index := 0; index < len(a); index++ {
		a[index] = rand.Intn(200)
	}
	/*
		fmt.Println(a)
		core(a)
		fmt.Println(a)
	*/
	fmt.Println(a)
	fmt.Println(mergeCore(a))
}

func Partition(s []int, lo, hi int) int {
	i, j := lo, hi
	x := s[lo]
	for i < j {
		for i < j {
			if s[j] >= x {
				j--
			} else {
				s[i] = s[j]
				break
			}
		}

		for i < j {
			if s[i] <= x {
				i++
			} else {
				s[j] = s[i]
				break
			}
		}
	}
	s[i] = x
	return i
}
func Partition_1(s []int, lo, hi int) int {
	x := lo + (hi-lo)/2
	small := lo - 1
	s[x], s[hi] = s[hi], s[x]
	for i := lo; i < hi; i++ {
		if s[i] <= s[hi] {
			small++
			if small != i {
				s[small], s[i] = s[i], s[small]
			}
		}
	}
	small++
	s[small], s[hi] = s[hi], s[small]
	return small
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	index := Partition_1(a, lo, hi)
	if index-1 > lo {
		quickSort(a, lo, index-1)
	}
	if index+1 < hi {
		quickSort(a, index+1, hi)
	}
}

func core(a []int) {
	lo, hi := 0, len(a)-1
	quickSort(a, lo, hi)
}

func merge(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	if i < len(a) {
		res = append(res, a[i:]...)
	} else if j < len(b) {
		res = append(res, b[j:]...)
	}
	return res
}

func mergeCore(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	left := mergeCore(a[0:mid])
	right := mergeCore(a[mid:])
	return merge(left, right)

}
