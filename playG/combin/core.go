package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 3, 4}
	fmt.Println(combine(a, 2))
	fmt.Println(com(a, 2))
}

func backTrace(a []int, start, k int, trace *[]int, res *[][]int) {
	if len(*trace) == k {
		r := make([]int, k)
		copy(r, (*trace))
		*res = append(*res, r)
		return
	}
	for i := start; i < len(a); i++ {
		*trace = append(*trace, a[i])
		backTrace(a, i+1, k, trace, res)
		*trace = (*trace)[0 : len(*trace)-1]
	}
}

func combine(a []int, k int) [][]int {
	if len(a) == 0 {
		return [][]int{{}}
	}
	res := [][]int{}
	trace := make([]int, 0, len(a))
	backTrace(a, 0, k, &trace, &res)
	return res
}
func com(a []int, k int) [][]int {
	res := [][]int{}
	var backT func(a []int, k, start int)
	trace := make([]int, 0)
	backT = func(a []int, k, start int) {
		if len(trace) == k {
			tmp := make([]int, len(trace))
			copy(tmp, trace)
			res = append(res, tmp)
		}
		for i := start; i < len(a); i++ {
			trace = append(trace, a[i])
			backT(a, k, i+1)
			trace = trace[0 : len(trace)-1]
		}
	}
	backT(a, k, 0)
	return res
}
