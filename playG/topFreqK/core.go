package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

type Node struct {
	val  int
	freq int
}
type Nodes []*Node

func (N Nodes) Len() int {
	return len(N)
}
func (N Nodes) Swap(a, b int) {
	N[a], N[b] = N[b], N[a]
}
func (N Nodes) Less(a, b int) bool {
	return N[a].freq < N[b].freq
}
func (N *Nodes) Push(a interface{}) {
	*N = append(*N, a.(*Node))
}
func (N *Nodes) Pop() interface{} {
	old := *N
	n := len(old)
	x := old[n-1]
	*N = old[0 : n-1]
	return x
}
func core(a []int, k int) []int {
	m := make(map[int]int)
	for _, item := range a {
		if m[item] == 0 {
			m[item] = 1
		} else {
			m[item]++
		}
	}
	ns := &Nodes{}
	i := 0
	for v, f := range m {
		if i < k {
			heap.Push(ns, &Node{
				val:  v,
				freq: f,
			})
		} else {
			if f > (*ns)[0].freq {
				(*ns)[0] = &Node{
					val:  v,
					freq: f,
				}
				heap.Fix(ns, 0)
			}
		}
		i++
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(ns).(*Node).val
	}
	return res
}
