package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := "ADOBECODEBANC"
	b := "ABC"
	fmt.Println(minWindow(a, b))
}
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; ok {
			m[t[i]]++
		} else {
			m[t[i]] = 1
		}
	}
	var i, e int
	count := len(t)
	var res string
	min := -1
	for ; e < len(s); e++ {
		if _, ok := m[s[e]]; ok {
			m[s[e]]--
			if m[s[e]] >= 0 {
				count--
			}
		}
		if count == 0 {
			for ; i <= e; i++ {
				if _, ok := m[s[i]]; ok {
					m[s[i]]++
					if m[s[i]] > 0 {
						count++
						break
					}
				}
			}
			if i <= e {
				length := e - i + 1
				if length < min || min < 0 {
					min = length
					res = string(s[i : e+1])
				}
			}
			i++
		}

	}
	return res
}
