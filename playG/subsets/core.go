package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 3, 4}
	res := [][]int{{}}
	subsets(a, &res)
	fmt.Println(res)
	fmt.Println(a)
	fmt.Println(subsets2(a))
}
func subsets(a []int, res *[][]int) {
	if len(a) == 0 {
		return
	}
	n := a[len(a)-1]
	b := a[0 : len(a)-1]
	subsets(b, res)
	length := len(*res)
	for i := 0; i < length; i++ {
		*res = append(*res, (*res)[i])
		//(*res)[i] = append((*res)[i], n)
		(*res)[len(*res)-1] = append((*res)[len(*res)-1], n)
	}
	return
}

func backTrace(a []int, start int, trace *[]int, res *[][]int) {
	r := make([]int, len(*trace))
	copy(r, (*trace))
	(*res) = append((*res), r)
	for i := start; i < len(a); i++ {
		(*trace) = append((*trace), a[i])
		backTrace(a, i+1, trace, res)
		(*trace) = (*trace)[0 : len(*trace)-1]
	}
}

func subsets2(a []int) [][]int {
	if len(a) == 0 {
		return [][]int{{}}
	}
	res := [][]int{}
	trace := make([]int, 0, len(a))
	backTrace(a, 0, &trace, &res)
	return res
}
