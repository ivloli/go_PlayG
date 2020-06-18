package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{3, 2, 1, 1, 2, 2, 6, 6, 4}
	fmt.Println(findKthLargest(a, 3))
}
func findKthLargest(nums []int, k int) int {
	return findCore(nums, 0, len(nums)-1, k)
}
func part(nums []int, s, e int) int {
	key := nums[s]
	i, j := s, e
	for i < j {
		for i < j {
			if nums[j] <= key {
				j--
			} else {
				nums[i] = nums[j]
				break
			}
		}
		for i < j {
			if nums[i] > key {
				i++
			} else {
				nums[j] = nums[i]
				break
			}
		}
	}
	nums[i] = key
	fmt.Println(key)
	fmt.Println(i)
	fmt.Println(nums)
	return i
}

func findCore(nums []int, s, e int, k int) int {
	pos := part(nums, s, e)
	if pos-s == k-1 {
		return nums[pos]
	} else if pos-s < k-1 {
		return findCore(nums, pos+1, e, k-(pos-s+1))
	} else {
		return findCore(nums, s, pos-1, k)
	}
}
