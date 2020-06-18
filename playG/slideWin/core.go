package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	a := []int{2, 3, 4, 2, 6, 2, 5, 1}
	fmt.Println(maxSlidingWindow(a, 3))
}

func maxSlidingWindow(a []int, k int) []int {
	res := make([]int, 0, len(a)-k+1)
	window := list.New()
	for index, value := range a {
		if index >= k && window.Front().Value.(int) <= index-k {
			window.Remove(window.Front())
		}
		for window.Len() != 0 && a[window.Back().Value.(int)] <= value {
			window.Remove(window.Back())
		}
		window.PushBack(index)
		if index >= k-1 {
			res = append(res, a[window.Front().Value.(int)])
		}
	}
	return res
}

type maxQ struct {
	Q *list.List
	M *list.List
}
type entry struct {
	data int
}

func NewMaxQ() *maxQ {
	m := maxQ{}
	m.Q = list.New()
	m.M = list.New()
	return &m
}

func (m *maxQ) Push(a int) {
	for m.M.Len() != 0 && a >= m.M.Back().Value.(*entry).data {
		m.M.Remove(m.M.Back())
	}
	e := &entry{a}
	m.M.PushBack(e)
	m.Q.PushBack(e)
}
func (m *maxQ) Pop() int {
	if m.Q.Len() == 0 {
		return 0
	}
	ele := m.Q.Front()
	if ele.Value.(*entry) == m.M.Front().Value.(*entry) {
		m.M.Remove(m.M.Front())
	}
	m.Q.Remove(ele)
	return ele.Value.(*entry).data
}
func (m *maxQ) GetMax() int {
	if m.M.Len() == 0 {
		return 0
	}
	return m.M.Front().Value.(*entry).data
}
