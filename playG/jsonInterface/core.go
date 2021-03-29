package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type resIf interface {
	GetName() string
}

type message struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Msg  string `json:"message"`
}
type msg struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Msg  string `json:"msg"`
}

func (m message) GetName() string {
	return m.Name
}
func (m msg) GetName() string {
	return m.Name
}

const F int = 2

func main() {
	fmt.Println("vim-go")
	var res resIf
	s1 := `{"name":"xiaoming","age":18,"message":"std"}`
	s2 := `{"name":"xiaoming","age":18,"msg":"std"}`
	if F == 1 {
		re := message{}
		jsoniter.Unmarshal([]byte(s1), &re)
		res = re
	} else {
		re := msg{}
		jsoniter.Unmarshal([]byte(s2), &re)
		res = re
	}
	fmt.Println(res)
}
