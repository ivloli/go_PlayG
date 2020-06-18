package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Node struct {
	next *Node
	data int
}

func revert(head *Node) *Node {
	if head == nil {
		return nil
	}
	var pre, cur, aft *Node
	cur = head
	for cur != nil {
		aft = cur.next
		cur.next = pre
		pre = cur

		cur = aft
	}
	return pre
}
func revert_part(head, tail *Node) *Node {
	if head == nil {
		return nil
	}
	var pre, cur, aft *Node
	cur = head
	for cur != tail {
		aft = cur.next
		cur.next = pre
		pre = cur
		cur = aft
	}
	return pre
}
func revertKGroup(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.next
	}
	newNode := revert_part(a, b)
	a.next = revertKGroup(b, k)
	return newNode
}

func isCross(h1, h2 *Node) bool {
	if h1 == nil || h2 == nil {
		return false
	}
	p1 := h1
	p2 := h2
	var f bool
	for {
		if p1 != nil {
			p1 = p1.next
		} else {
			if f == false {
				p1 = h2
				f = true
			} else {
				return false
			}
		}
		if p2 != nil {
			p2 = p2.next
		} else {
			p2 = h1
		}
		if p1 == p2 {
			return true
		}
	}
}

func Loop(head *Node) (bool, *Node) {
	if head == nil {
		return false, nil
	}
	p1, p2 := head, head

	for p2 != nil && p2.next != nil {
		p1 = p1.next
		p2 = p2.next.next
		if p1 == p2 {
			p := p1
			p2 = head.next
			p1 = p1.next
			for p1 != p {
				p1 = p1.next
				p2 = p2.next
			}
			p1 = head
			for p1 != p2 {
				p1 = p1.next
				p2 = p2.next
			}
			return true, p1
		}
	}
	return false, nil
}
