package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := "AaBbCcDd"
	fmt.Println(core(a))
}

func core(s string) string {
	if len(s) < 2 {
		return s
	}
	b := []byte(s)
	p := len(b) - 1
	for p >= 0 {
		if b[p] >= 'A' && b[p] <= 'Z' {
			i := p
			for i < len(b)-1 {
				if b[i+1] < 'A' || b[i+1] > 'Z' {
					b[i+1], b[i] = b[i], b[i+1]
				}
				i++
			}
		}
		p--
	}
	return string(b)
}
