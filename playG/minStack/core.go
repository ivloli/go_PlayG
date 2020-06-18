package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	a := []int{2, 3, 4, 5, 6, 3, 3, 4, 5, 4, 3, 1, 0, 0, 1}
	fmt.Println(a)
	ms := NewMinStack()
	for _, item := range a {
		ms.Push(item)
	}
	for ms.GetLen() != 0 {
		fmt.Println(ms.GetMin(), "|", ms.Pop())
	}
}

type MinStack struct {
	list *list.List
	min  int
}

func NewMinStack() MinStack {
	m := MinStack{}
	m.list = list.New()
	return m
}

func (m *MinStack) Push(a int) {
	if m.list.Len() == 0 {
		m.min = a
		m.list.PushBack(0)
	} else {
		m.list.PushBack(a - m.min)
		if a < m.min {
			m.min = a
		}
	}
}

func (m MinStack) GetMin() int {
	return m.min
}
func (m MinStack) GetLen() int {
	return m.list.Len()
}
func (m *MinStack) Pop() int {
	if m.list.Len() == 0 {
		return 0
	}
	ele := m.list.Back()
	n := ele.Value.(int)
	m.list.Remove(ele)
	if n >= 0 {
		return m.min + n
	}
	oldMin := m.min
	m.min = oldMin - n
	return oldMin
}
