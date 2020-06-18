package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Node struct {
	Data        int
	Left, Right *Node
}

func isBST(root *Node, pLast **Node) bool {
	if root == nil {
		return true
	}
	b := isBST(root.Left, pLast)
	if b && ((*pLast) == nil || (*pLast).Data <= root.Data) {
		*pLast = root
		return isBST(root.Right, pLast)
	} else {
		return false
	}
}

func IsBST(root *Node) bool {
	var last *Node
	return isBST(root, &last)
}
