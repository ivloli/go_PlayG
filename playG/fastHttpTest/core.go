package main

import (
	"core/visit"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	reqPara := visit.ReqParmater{
		Method:      "POST",
		Url:         "http://127.0.0.1:18080/tizi365",
		ContentType: "application/json",
		Body:        []byte("pak=13"),
	}
	r, e := visit.DoRequest(&reqPara, nil)
	fmt.Println(e)
	fmt.Println(string(r))
}
