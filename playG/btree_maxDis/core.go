package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
func getDisDepth(root *Node) (dep, dis int) {
	if root == nil {
		return
	}
	ldep, ldis := getDisDepth(root.left)
	rdep, rdis := getDisDepth(root.right)
	dep = max(ldep, rdep) + 1
	dis = max(max(ldis, rdis), ldep+rdep)
	return
}
func GetMaxDis(root *Node) int {
	_, dis := getDisDepth(root)
	return dis
}
