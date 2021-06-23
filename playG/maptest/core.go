package main

import (
	"fmt"
	"strings"
)

var tmap = map[int]bool{
	1: true,
	2: true,
}

func main() {
	fmt.Println("vim-go")
	mmap := make(map[string]*User, 0)
	mslice := make([]User, 0)
	mslice = append(mslice, User{"xiaoli", 2}, User{"xiaohong", 3})
	for index, item := range mslice {
		mmap[item.Name] = &(mslice[index])
	}
	sstr := []string{"xiaoli", "xiaohong"}
	for _, item := range sstr {
		if val, ok := mmap[item]; ok {
			val.Age++
		}
	}
	for val, value := range mslice {
		fmt.Println(val)
		fmt.Println(value)
	}
	fmt.Println(len(strings.Split("", ",")))
	if !tmap[1] {
		fmt.Println("!")
	} else {
		fmt.Println("?")
	}
	if !tmap[3] {
		fmt.Println("!")
	} else {
		fmt.Println("?")
	}
}

type User struct {
	Name string
	Age  int
}
