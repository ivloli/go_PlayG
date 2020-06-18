package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func core(n, k int) int {
	if n == 0 {
		return 0
	}
	same, diff := 0, k
	for i := 2; i <= n; i++ {
		temp := diff
		diff = (same + diff) * (k - 1)
		same = temp
	}
	return same + diff
}
