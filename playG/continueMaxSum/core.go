package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func core(a []int) int {
	dp := make([]int, len(a))
	dp[0] = a[0]
	res := dp[0]
	for i := 1; i < len(a); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + a[i]
		} else {
			dp[i] = a[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
