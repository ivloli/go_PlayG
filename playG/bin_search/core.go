package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 3, 3, 3, 3, 4, 7, 8, 9, 10}
	fmt.Println(LeftBound2(a, 3))
	fmt.Println(RightBound2(a, 3))
	fmt.Println(LeftBound2(a, 5))
	fmt.Println(RightBound2(a, 5))
}

func core(a []int, t int) int {
	if len(a) == 0 {
		return -1
	}
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == t {
			return mid
		} else if a[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func LeftBound(a []int, t int) int {
	if len(a) == 0 {
		return -1
	}
	left, right := 0, len(a)
	for left < right {
		mid := left + (right-left)/2
		if a[mid] == t {
			right = mid
		} else if a[mid] > t {
			right = mid
		} else if a[mid] < t {
			left = mid + 1
		}
	}
	if left == len(a) {
		return -1
	}
	if a[left] != t {
		return -1
	} else {
		return left
	}
}

func LeftBound2(a []int, t int) (bool, int) {
	if len(a) == 0 {
		return false, -1
	}
	left, right := 0, len(a)
	for left < right {
		mid := left + (right-left)/2
		if a[mid] == t {
			right = mid
		} else if a[mid] > t {
			right = mid
		} else if a[mid] < t {
			left = mid + 1
		}
	}
	if left == len(a) {
		return false, left - 1
	}
	if a[left] != t {
		if left == 0 {
			return false, 0
		} else {
			return false, left - 1
		}
	} else {
		return true, left
	}
}

func RightBound(a []int, t int) int {
	if len(a) == 0 {
		return -1
	}
	left, right := 0, len(a)
	for left < right {
		mid := left + (right-left)/2
		if a[mid] == t {
			left = mid + 1
		} else if a[mid] < t {
			left = mid + 1
		} else if a[mid] > t {
			right = mid
		}
	}
	if right == 0 {
		return -1
	}
	if a[right-1] != t {
		return -1
	} else {
		return right - 1
	}
}

func RightBound2(a []int, t int) (bool, int) {
	if len(a) == 0 {
		return false, -1
	}
	left, right := 0, len(a)
	for left < right {
		mid := left + (right-left)/2
		if a[mid] == t {
			left = mid + 1
		} else if a[mid] < t {
			left = mid + 1
		} else if a[mid] > t {
			right = mid
		}
	}
	if right == 0 {
		return false, 0
	}
	if a[right-1] != t {
		if right == len(a) {
			return false, right - 1
		} else {
			return false, right
		}
	} else {
		return true, right - 1
	}
}
