package main

import "fmt"

func main() {
	a := "abcdefghijkabcdefzvxwel12345o7896abc"
	fmt.Println(a)
	fmt.Println(core(a))

}
func core(s string) (int, string) {
	mmap := make(map[byte]int, 256)
	var maxlen int
	var start, end int
	var res string
	for index, item := range []byte(s) {
		end = index
		if val, ok := mmap[item]; ok {
			if val >= start {
				curlen := end - start
				if curlen > maxlen {
					maxlen = curlen
					res = s[start:end]
					fmt.Println(res)
				}
				start = val + 1
			}
		}
		mmap[item] = index
	}
	curlen := len(s) - start
	if curlen > maxlen {
		maxlen = curlen
		res = s[start:len(s)]
	}
	return maxlen, res
}
