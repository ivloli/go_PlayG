package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type set struct {
	Len int
	m   map[interface{}]bool
}

func NewSet() set {
	s := set{}
	s.Len = 0
	s.m = make(map[interface{}]bool)
	return s
}

func (s *set) Add(i interface{}) {
	s.m[i] = true
	s.Len = len(s.m)
}

func (s *set) IsEx(i interface{}) bool {
	_, ok := s.m[i]
	return ok
}

func (s *set) Remove(i interface{}) {
	if _, ok := s.m[i]; ok {
		delete(s.m, i)
	}
}

func (s *set) Size() int {
	return s.Len
}
