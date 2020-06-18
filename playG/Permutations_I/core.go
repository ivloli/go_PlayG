package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	/*
		a := "abc"
		res := core(a)
		for _, item := range res {
			fmt.Println(item)
		}
		b := "abc"
		res = core_duplicate(b)

		for _, item := range res {
			fmt.Println(item)
		}
	*/
	c := "abc"
	resc := backTreeCore(c, 1)
	for _, item := range resc {
		fmt.Println(item)
	}
}

func core(s string) []string {
	b := []byte(s)
	res := make([]string, 0)
	DFS(b, 0, &res)
	return res
}

func core_duplicate(s string) []string {
	res := make([]string, 0)
	DFS_duplicate(s, 0, &res)
	return res
}
func DFS(b []byte, start int, res *[]string) {
	if start >= len(b) {
		(*res) = append((*res), string(b))
	}
	for i := start; i < len(b); i++ {
		if start != i {
			b[start], b[i] = b[i], b[start]
		}
		DFS(b, start+1, res)
		if start != i {
			b[start], b[i] = b[i], b[start]
		}
	}
}

func DFS_duplicate(s string, start int, res *[]string) {
	if start >= len(s) {
		(*res) = append((*res), s)
	}
	b := []byte(s)
	for i := start; i < len(b); i++ {
		if i != start && b[start] == b[i] {
			continue
		}
		b[start], b[i] = b[i], b[start]
		DFS_duplicate(string(b), start+1, res)
	}
}

func backTreeCore(s string, k int) []string {
	res := make([]string, 0)
	used := make([]bool, len(s))
	track := make([]byte, 0, len(s))
	backTree([]byte(s), k, &used, &track, &res)
	return res
}
func backTree(nums []byte, k int, used *[]bool, track *[]byte, res *[]string) {
	if len(*track) == k {
		*res = append(*res, string(*track))
		return
	}
	for i := 0; i < len(nums); i++ {
		if (*used)[i] == false {
			(*used)[i] = true
			*track = append(*track, nums[i])
			backTree(nums, k, used, track, res)
			*track = (*track)[0 : len(*track)-1]
			(*used)[i] = false
		}
	}
}
