package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "xiaoqiang:+86 12345678"
	b := "李克强：+86 123456777"
	fmt.Printf(core(a))
	fmt.Printf(core(b))
}

func core(s string) string {
	res := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsDigit(r) {
			res = append(res, r)
		}
	}
	return string(res)
}
