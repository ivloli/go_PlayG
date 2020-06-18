package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 4, 5, 3, 6, 7}
	b := []int{4, 2, 5, 6, 1, 3, 7}
	c := buildTree(a, b)
	fmt.Println(c)
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func buildTree(p, in []int) *Node {
	m := make(map[int]int, len(in))
	for index, item := range in {
		m[item] = index
	}
	ok, tree := build(p, m, 0)
	if ok {
		return tree
	} else {
		return nil
	}
}

func build(p []int, m map[int]int, offset int) (bool, *Node) {
	if len(p) == 0 {
		return true, nil
	}
	root := p[0]
	idx := m[root] - offset
	if idx < 0 || 1+idx >= len(p) {
		return false, nil
	}
	ok1, l := build(p[1:1+idx], m, offset)
	ok2, r := build(p[1+idx:], m, offset+idx+1)
	if ok1 && ok2 {
		return true, &Node{
			data:  root,
			left:  l,
			right: r,
		}
	} else {
		return false, nil
	}
}
