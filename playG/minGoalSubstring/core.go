package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := "ABCDEEFFGGTSSAAA"
	b := "GGEFE"
	fmt.Println(core(a, b))
}

func core(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	tmap := make(map[byte]int)
	smap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}
	l, cnt, min := 0, 0, -1
	res := ""

	for r := 0; r < len(s); r++ {
		if tmap[s[r]] == 0 {
			continue
		}
		if smap[s[r]] < tmap[s[r]] {
			cnt++
		}
		smap[s[r]]++
		if cnt == len(t) {
			k := s[l]
			for tmap[k] == 0 || smap[k] > tmap[k] {
				if smap[k] > tmap[k] && tmap[k] > 0 {
					smap[k]--
				}
				l++
				k = s[l]
			}
			if min == -1 || r-l+1 < min {
				min = r - l + 1
				res = s[l : r+1]
			}
		}
	}
	return res
}
