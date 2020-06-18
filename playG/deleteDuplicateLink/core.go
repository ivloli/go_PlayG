package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{2, 3, 4, 5, 5, 5, 6, 6, 7, 6, 7, 8}
	fmt.Println(a)
	h := &Node{a[0], nil}
	c := h
	for i := 1; i < len(a); i++ {
		c.Next = &Node{a[i], nil}
		c = c.Next
	}
	h1 := deleteDuplicate(h)

	for h1 != nil {
		fmt.Println(h1.Data)
		h1 = h1.Next
	}
}

type Node struct {
	Data int
	Next *Node
}

func deleteDuplicate(h *Node) *Node {
	if h == nil || h.Next == nil {
		return h
	}
	if h.Data == h.Next.Data {
		h = h.Next
		return deleteDuplicate(h)
	}
	h.Next = deleteDuplicate(h.Next)
	return h
}
