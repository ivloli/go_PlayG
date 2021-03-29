package main

import "fmt"

type A struct {
	Name string
}
type AA struct {
	*A
}

func (s *A) GetName() string {
	return s.Name
}

func (s *AA) GetName() string {
	return s.Name + ":NEW"
}
func main() {
	fmt.Println("vim-go")
	s1 := AA{
		A: &A{
			Name: "xiaoming",
		},
	}
	fmt.Println(s1.GetName())
}
