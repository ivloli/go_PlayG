package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	a := "a{b,c}d,e{x,y}j,k"
	fmt.Println(core(a))
}
func core(a string) []string {
	temp := [][]int{}
	for i := 0; i < len(a); i++ {
		if a[i] == '{' {
			temp = append(temp, []int{i})
		} else if a[i] == '}' {
			temp[len(temp)-1] = append(temp[len(temp)-1], i)
		}
	}
	strs := [][]string{}
	for _, item := range temp {
		if item[1]-item[0] > 2 {
			segs := strings.Split(string(a[item[0]+1:item[1]]), ",")
			substrs := []string{}
			for _, v := range segs {
				if v != "" {
					substrs = append(substrs, v)
				}
			}
			strs = append(strs, substrs)
		}

	}
	res := []string{""}
	j := 0
	for i := 0; i < len(a); {
		if j == len(strs) || i < temp[j][0] {
			if a[i] != ',' {
				for index, _ := range res {
					res[index] = res[index] + string(a[i])
				}
			}
			i++
		} else if i == temp[j][0] {
			i = temp[j][1]
			length := len(res)
			for k := 0; k < length; k++ {
				for n := 1; n < len(strs[j]); n++ {
					t := res[k] + string(strs[j][n])
					res = append(res, t)
				}
			}
			for k := 0; k < length; k++ {
				res[k] = res[k] + strs[j][0]
			}
			i++
			if j < len(strs) {
				j++
			}
		}
	}
	result := []string{}
	for _, item := range res {
		result = append(result, string(item))
	}
	return result
}
