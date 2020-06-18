package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func core(a []int) int {
	if len(a) < 3 {
		return 0
	}
	res := 0
	lm, rm := a[0], a[len(a)-1]
	l, r := 0, len(a)-1
	for l <= r {
		lm = max(lm, a[l])
		rm = max(rm, a[r])
		if lm < rm {
			res += (lm - a[l])
			l++
		} else {
			res += (rm - a[r])
			r--
		}
	}
	return res
}

func core_1(a []int) int {
	if len(a) < 3 {
		return 0
	}
	res := 0
	var lm, rm int
	l, r := 0, len(a)-1
	for l < r {
		if lm <= rm {
			for l < r {
				if a[l] <= lm {
					res += (lm - a[l])
				} else {
					lm = a[l]
					if lm > rm {
						break
					}
				}
				l++
			}
		} else {
			for l < r {
				if a[r] <= rm {
					res += (rm - a[r])
				} else {
					rm = a[r]
					if rm > lm {
						break
					}
				}
				r--
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
