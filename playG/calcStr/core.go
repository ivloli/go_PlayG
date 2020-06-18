package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	a := "1+2*3+(3+(5+4)*2)"
	fmt.Println(core(a))
}

func calc(a, b int, op byte) int {
	var res int
	switch op {
	case '*':
		res = a * b
	case '/':
		res = a / b
	case '+':
		res = a + b
	case '-':
		res = a - b
	}
	return res
}

func calcSimpleStr(s []byte) int {
	fmt.Println(string(s))
	nums := []int{}
	ops := []byte{}
	cur := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			cur = append(cur, s[i])
		} else {
			num, _ := strconv.Atoi(string(cur))
			nums = append(nums, num)
			ops = append(ops, s[i])
			i++
			cur = []byte{s[i]}
		}
	}
	num, _ := strconv.Atoi(string(cur))
	nums = append(nums, num)
	s1 := []int{nums[0]}
	s2 := []byte{}
	for i, j := 1, 0; i < len(nums); i, j = i+1, j+1 {
		if ops[j] == '+' || ops[j] == '-' {
			s2 = append(s2, ops[j])
			s1 = append(s1, nums[i])
		} else {
			a := calc(s1[len(s1)-1], nums[i], ops[j])
			s1[len(s1)-1] = a
		}
	}
	var a int
	a = s1[0]
	for i := 1; i < len(s1); i++ {
		a = calc(a, s1[i], s2[i-1])
	}
	return a
}

func remove(s string) string {
	b := []byte(s)
	st := []int{}
	for i := 0; i < len(b); i++ {
		if b[i] == '(' {
			st = append(st, i)
		} else if b[i] == ')' {
			s := st[len(st)-1]
			st = st[0 : len(st)-1]
			e := i
			r := calcSimpleStr(b[s+1 : e])
			rb := []byte(strconv.Itoa(r))
			if e+1 < len(b) {
				b = append(b[0:s], append(rb, b[e+1:]...)...)
			} else {
				b = append(b[0:s], rb...)
			}
			i = s - 1 + len(rb)
		}
	}
	return string(b)
}

func core(s string) int {
	s1 := remove(s)
	return calcSimpleStr([]byte(s1))
}
