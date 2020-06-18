package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

type IntHeap struct {
	Data []float64
	Dir  bool
}

func (h IntHeap) Len() int {
	return len(h.Data)
}

func (h IntHeap) Swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func (h IntHeap) Less(i, j int) bool {
	if h.Dir {
		return h.Data[i] < h.Data[j]
	} else {
		return h.Data[i] > h.Data[j]
	}
}

func (h IntHeap) BelowTop(a float64) bool {
	if h.Len() == 0 {
		return false
	}
	if h.Dir {
		return a > h.Data[0]
	} else {
		return a < h.Data[0]
	}
}

func (h *IntHeap) Push(i interface{}) {
	(*h).Data = append((*h).Data, i.(float64))
}

func (h *IntHeap) Pop() interface{} {
	old := (*h).Data
	(*h).Data = old[0 : len(old)-1]
	x := old[len(old)-1]
	return x
}

type Solution struct {
	minHeap *IntHeap
	maxHeap *IntHeap
}

func NewSolution() Solution {
	s := Solution{}
	s.maxHeap = &IntHeap{
		Data: []float64{},
		Dir:  false,
	}
	s.minHeap = &IntHeap{
		Data: []float64{},
		Dir:  true,
	}
	return s
}

func (s *Solution) Insert(a float64) {
	if s.maxHeap == nil || s.minHeap == nil {
		s.maxHeap = &IntHeap{
			Data: []float64{},
			Dir:  false,
		}
		s.minHeap = &IntHeap{
			Data: []float64{},
			Dir:  true,
		}
	}
	h := s.maxHeap
	if s.minHeap.Len() == s.maxHeap.Len() {
		if s.minHeap.Len() != 0 && (*s.minHeap).Data[0] < a {
			h = s.minHeap
		}
		heap.Push(h, a)
		return
	}
	if s.minHeap.Len() < s.maxHeap.Len() {
		h = s.minHeap
	}
	h1 := s.minHeap
	if h == s.minHeap {
		h1 = s.maxHeap
	}
	if h1.BelowTop(a) {
		h1.Data[0], a = a, h1.Data[0]
		heap.Fix(h1, 0)
	}
	heap.Push(h, a)
}

func (s *Solution) GetMedian() (bool, float64) {
	if s.minHeap.Len() == s.maxHeap.Len() {
		if s.minHeap.Len() == 0 {
			return false, 0
		}
		return true, s.maxHeap.Data[0] + (s.minHeap.Data[0]-s.maxHeap.Data[0])/2
	}
	if s.maxHeap.Len() > s.minHeap.Len() {
		return true, s.maxHeap.Data[0]
	} else {
		return true, s.minHeap.Data[0]
	}
}
