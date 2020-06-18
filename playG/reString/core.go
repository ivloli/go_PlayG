package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := "we are the one"
	fmt.Println(core(a))
}

func rev(a []byte, s, e int) {
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}
}

func core(a string) string {
	b := []byte(a)
	s := 0
	inStr := false
	for i := 0; i < len(b); i++ {
		if inStr {
			if b[i] == ' ' {
				rev(b, s, i-1)
				inStr = false
			} else {
				continue
			}
		} else {
			if b[i] == ' ' {
				continue
			} else {
				s = i
				inStr = true
			}
		}
	}
	if inStr {
		rev(b, s, len(b)-1)
	}
	rev(b, 0, len(b)-1)
	return string(b)
}
