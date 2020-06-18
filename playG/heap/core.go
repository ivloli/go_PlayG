package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type heap []int

func (h heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h heap) fixup(i int) {
	fa := (i - 1) / 2
	for fa >= 0 && h.less(i, fa) {
		h.swap(fa, i)
		i = fa
		fa = (i - 1) / 2
	}
}

func (h heap) fixdown(i int) {
	l := len(h)
	s := 2 * (i + 1)
	for s < l {
		if s+1 < l && h.less(s+1, s) {
			s++
		}
		if h.less(s, i) {
			h.swap(s, i)
			i = s
			s = 2 * (i + 1)
		} else {
			break
		}
	}
}

func (h *heap) Add(i int) {
	if h == nil {
		return
	}
	l := len(*h)
	*h = append(*h, i)
	h.fixup(l)
}

func (h *heap) Pop() (bool, int) {
	if h == nil || len(*h) == 0 {
		return false, 0
	}
	res := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	x := (*h)[0 : len(*h)-1]
	*h = x
	h.fixdown(0)
	return true, res
}

func (h heap) Replace(i int) (bool, int) {
	if len(h) == 0 {
		return false, 0
	}
	res := h[0]
	h[0] = i
	h.fixdown(0)
	return true, res
}

func (h heap) Init() {
	if len(h) < 2 {
		return
	}
	for i := (len(h) / 2) - 1; i >= 0; i-- {
		h.fixdown(i)
	}
}
