package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type users struct {
	Data []user `json:"data"`
}

func (self *users) Show() {
	fmt.Println(self)
	fmt.Println(self.Data)
	b, _ := json.Marshal(self)
	fmt.Println(string(b))
	fmt.Println(len(self.Data))
}
func main() {
	ustr := `{"data":[{"name":"xiaoliu","age":10},{"name":"xiaoli","age":11}]}`
	var sts users
	json.Unmarshal([]byte(ustr), &sts)
	sts.Show()
}
