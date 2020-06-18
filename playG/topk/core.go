package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	a := []int{10, 12, 33, 44, 13, 15, 16, 36, 78, 77, 89, 76, 55, 67, 43, 12, 101, 103, 103, 104, 106, 109, 106}
	b := a[0:10]
	c := IntHeap(b)
	h := &c
	heap.Init(h)
	for _, v := range a[10:len(a)] {
		if v > (*h)[0] {
			(*h)[0] = v
			heap.Fix(h, 0)
		}
	}
	fmt.Println(*h)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Insert(a int) {
	if h.Len() < 10 {
		heap.Push(h, a)
		heap.Fix(h, 0)
	} else {
		if a > (*h)[0] {
			(*h)[0] = a
			heap.Fix(h, 0)
		}
	}
}

func (h *IntHeap) GetTopK() []int {
	return []int(*h)
}
