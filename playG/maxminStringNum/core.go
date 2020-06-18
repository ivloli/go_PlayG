package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	a := []int{9, 10, 23, 45, 10, 0, 0, 234}
	fmt.Println(core(a))
}

func Comp(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			return true
		} else if a[i] > b[j] {
			return false
		}
		i++
		j++
	}
	if len(a) == len(b) {
		return true
	}
	if len(a) < len(b) {
		return Comp(a, b[j:])
	} else {
		return Comp(a[i:], b)
	}
}

type nums []int

func (n nums) Len() int {
	return len(n)
}

func (n nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nums) Less(i, j int) bool {
	a := strconv.Itoa(n[i])
	b := strconv.Itoa(n[j])
	return !Comp(a, b)
}

func core(a []int) string {
	n := nums(a)
	sort.Sort(n)
	sb := strings.Builder{}
	firstZero := true
	for _, v := range n {
		if firstZero {
			if v == 0 {
				continue
			}
			firstZero = false
		}
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}
