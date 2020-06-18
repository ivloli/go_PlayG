package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(a))
}
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}
	l, r := 1, 1
	max := len(nums) - 1
	for i := 1; i < len(nums); i++ {
		fmt.Println(l)
		res[i] = res[i] * l
	}
	for j := max - 1; j >= 0; j-- {
		r = r * nums[j+1]
		res[j] = res[j] * r
	}
	return res
}
