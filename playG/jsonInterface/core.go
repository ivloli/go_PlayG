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
	Flag bool   `json:"flag"`
}
type baseMsg struct {
	Name string
	Age  int
}
type newMsg struct {
	*baseMsg
}

func (m *baseMsg) GetName() string {
	return m.Name
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

	s3 := `{"name":"xiaoming","age":18,"msg":"std"}`

	res3 := msg{}
	jsoniter.Unmarshal([]byte(s3), &res3)

	fmt.Println(res3)

	nmsg := &newMsg{
		baseMsg: &baseMsg{
			Name: "basename",
			Age:  1,
		},
	}
	fmt.Printf("%+v", *(nmsg.baseMsg))
}
