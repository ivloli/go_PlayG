package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func core_judge(a []int) bool {
	if len(a) < 2 {
		return true
	}
	i, j := 0, len(a)-1
	for i < j {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func helper(a []int, i, j int) (int, []int) {
	for i >= 0 && j < len(a) {
		if a[i] == a[j] {
			i--
			j++
		}
	}
	length := j - i - 1
	return length, a[i+1 : j]
}

func core_find(a []int) (int, []int) {
	reslen := 0
	res := a[0:0]
	for i := 0; i < len(a); i++ {
		m1, s1 := helper(a, i, i)
		m2, s2 := helper(a, i, i+1)
		if m1 > reslen {
			if m1 > m2 {
				reslen = m1
				res = s1
			} else if m2 > reslen {
				reslen = m2
				res = s2
			}
		} else if m2 > reslen {
			reslen = m2
			res = s2
		}
	}
	return reslen, res
}
