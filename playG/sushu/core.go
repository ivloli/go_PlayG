package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(core(20))
}

func is(a int) bool {
	if a < 2 {
		return true
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
func core(a int) []int {
	nums := make([]bool, a+1)
	for i := 2; i*i <= a; i++ {
		if is(i) {
			for j := i * i; j <= a; j += i {
				nums[j] = true
			}
		}
	}
	res := make([]int, 0, a+1)
	for i := 1; i <= a; i++ {
		if nums[i] == false {
			res = append(res, i)
		}
	}
	return res
}
