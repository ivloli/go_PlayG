package main

import (
	"fmt"
	_ "sort"
)

func main() {
	fmt.Println("vim-go")
	a := []item{{1, 3}, {2, 4}, {3, 5}, {4, 6}, {7, 10}, {8, 25}}
	fmt.Println(oneZero(a, 15))
	fmt.Println(Complete(a, 15))
}

type item struct {
	w int
	v int
}
type items []item

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func oneZero(a []item, W int) int {
	if len(a) == 0 {
		return 0
	}
	dp := make([]int, W+1)

	for _, item := range a {
		for n := W; n-item.w >= 0; n-- {
			dp[n] = max(dp[n-item.w]+item.v, dp[n])
		}
		fmt.Println(dp)
	}
	return dp[W]
}

func (t items) Len() int {
	return len(t)
}
func (t items) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t items) Less(i, j int) bool {
	return t[i].w < t[j].w
}

func Complete(a []item, W int) int {
	if len(a) == 0 {
		return 0
	}
	dp := make([]int, W+1)

	for _, item := range a {
		for n := item.w; n <= W; n++ {
			dp[n] = max(dp[n], dp[n-item.w]+item.v)
		}
		fmt.Println(dp)
	}
	return dp[W]
}
