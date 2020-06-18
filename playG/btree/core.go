package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

type Node struct {
	Data        int
	Left, Right *Node
}

func FindPath(root *Node, t int) (bool, *list.List) {
	if root == nil {
		return false, nil
	}
	s := list.New()
	n := root
	var pre *Node
	for s.Len() != 0 || n != nil {
		for n != nil {
			s.PushBack(n)
			if n.Data == t {
				return true, s
			}
			n = n.Left
		}
		n = (s.Back()).Value.(*Node)
		for n.Right == nil || n.Right == pre {
			pre = n
			s.Remove(s.Back())
			if s.Len() == 0 {
				return false, nil
			}
			n = (s.Back()).Value.(*Node)
		}
		n = n.Right
	}
	return false, nil
}

func FindLCA(root *Node, t1, t2 int) (bool, *Node) {
	if root == nil {
		return false, nil
	}
	s1 := list.New()
	s2 := list.New()
	n := root
	var pre *Node
	var f1, f2 bool
	for s1.Len() != 0 || s2.Len() != 0 || n != nil {
		for n != nil {
			if f1 != true {
				s1.PushBack(n)
			}
			if f2 != true {
				s2.PushBack(n)
			}
			if n.Data == t1 {
				f1 = true
			}
			if n.Data == t2 {
				f2 = true
			}
			if f1 && f2 {
				break
			}
			n = n.Left
		}
		if f1 && f2 {
			break
		}
		if f1 != true {
			n = (s1.Back()).Value.(*Node)
		} else {
			n = (s2.Back()).Value.(*Node)
		}
		for n.Right == nil || n.Right == pre {
			pre = n
			if f1 != true {
				s1.Remove(s1.Back())
			}
			if f2 != true {
				s2.Remove(s2.Back())
			}
			if s1.Len() == 0 || s2.Len() == 0 {
				return false, nil
			}
			if f1 != true {
				n = (s1.Back()).Value.(*Node)
			} else {
				n = (s2.Back()).Value.(*Node)
			}
		}
		n = n.Right
	}
	if f1 && f2 {
		for s1.Len() != 0 && s2.Len() != 0 {
			c1 := s1.Back().Value.(*Node)
			c2 := s2.Back().Value.(*Node)
			if c1 == c2 {
				return true, c1
			}
			s1.Remove(s1.Back())
			s2.Remove(s2.Back())
		}
	}

	return false, nil
}

func PreVisit(root *Node) {
	s := list.New()
	for root != nil || s.Len() != 0 {
		for root != nil {
			fmt.Println(root.Data)
			s.PushBack(root)
			root = root.Left
		}
		root = s.Back().Value.(*Node)
		s.Remove(s.Back())
		root = root.Right
	}
}

func MidVisit(root *Node) {
	s := list.New()
	for root != nil || s.Len() != 0 {
		for root != nil {
			s.PushBack(root)
			root = root.Left
		}
		root = s.Back().Value.(*Node)
		fmt.Println(root.Data)
		s.Remove(s.Back())
		root = root.Right
	}
}

func PostVisit1(root *Node) {
	if root == nil {
		return
	}
	s1 := list.New()
	s2 := list.New()
	s1.PushBack(root)
	for s1.Len() != 0 {
		element := s1.Back()
		n := element.Value.(*Node)
		s2.PushBack(n)
		if n.Left != nil {
			s1.PushBack(n.Left)
		}
		if n.Right != nil {
			s1.PushBack(n.Right)
		}
	}
	for s2.Len() != 0 {
		element := s2.Back()
		fmt.Println(element.Value.(*Node).Data)
		s2.Remove(element)
	}
}

func PostVisit2(root *Node) {
	if root == nil {
		return
	}
	s := list.New()
	s.PushBack(root)
	var n, pre *Node
	for s.Len() != 0 {
		n = s.Back().Value.(*Node)
		if (n.Right == nil && n.Left == nil) || (pre != nil && (pre == n.Right || pre == n.Left)) {
			fmt.Println(n.Data)
			pre = n
		} else {
			if n.Right != nil {
				s.PushBack(n.Right)
			}
			if n.Left != nil {
				s.PushBack(n.Left)
			}
		}
	}
}

func LevelVisit(root *Node) {
	if root == nil {
		return
	}
	s := list.New()
	s.PushBack(root)
	var element *list.Element
	for s.Len() != 0 {
		element = s.Front()
		root = element.Value.(*Node)
		fmt.Println(root.Data)
		if root.Left != nil {
			s.PushBack(root.Left)
		}
		if root.Right != nil {
			s.PushBack(root.Right)
		}
		s.Remove(element)
	}
}

func LevelLineVisit(root *Node) {
	if root == nil {
		return
	}
	s := list.New()
	s.PushBack(root)
	nextLevel := 0
	curLevel := 1
	var element *list.Element
	for s.Len() != 0 {
		element = s.Front()
		root = element.Value.(*Node)
		fmt.Println(root.Data)
		if root.Left != nil {
			s.PushBack(root.Left)
			nextLevel++
		}
		if root.Right != nil {
			s.PushBack(root.Right)
			nextLevel++
		}
		s.Remove(element)
		curLevel--
		if curLevel == 0 {
			curLevel = nextLevel
			nextLevel = 0
			fmt.Println("")
		}
	}
}

func ZLevelVisit(root *Node) {
	if root == nil {
		return
	}
	s1 := list.New()
	s2 := list.New()
	var element *list.Element
	s1.PushBack(root)
	for s1.Len() != 0 || s2.Len() != 0 {
		for s1.Len() != 0 {
			element = s1.Back()
			fmt.Printf("%d ", element.Value.(*Node).Data)
			root = element.Value.(*Node)
			if root.Left != nil {
				s2.PushBack(root.Left)
			}
			if root.Right != nil {
				s2.PushBack(root.Right)
			}
		}
		for s2.Len() != 0 {
			element = s2.Back()
			fmt.Printf("%d ", element.Value.(*Node).Data)
			root = element.Value.(*Node)
			if root.Right != nil {
				s1.PushBack(root.Right)
			}
			if root.Left != nil {
				s1.PushBack(root.Left)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Maxdepth(root *Node) int {
	if root == nil {
		return 0
	}
	l := Maxdepth(root.Left)
	r := Maxdepth(root.Right)
	return max(l, r) + 1
}
func MinDepth(root *Node) int {
	if root == nil {
		return 0
	}
	l := MinDepth(root.Left)
	r := MinDepth(root.Right)
	if root.Left != nil && root.Right != nil {
		return min(l, r) + 1
	} else {
		return max(l, r) + 1
	}
}
