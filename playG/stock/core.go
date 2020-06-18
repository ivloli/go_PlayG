package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func k1(a []int) int {
	if len(a) == 0 {
		return 0
	}
	dp_0 := 0
	dp_1 := -1 * a[0]
	for _, v := range a[1:] {
		dp_0 = max(dp_0, dp_1+v)
		dp_1 = max(dp_1, 0-v)
	}
	return dp_0
}
