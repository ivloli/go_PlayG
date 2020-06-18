package main

import (
	"fmt"
	"time"
)

func main() {
	createTime := "2020-02-27 16:20:23"
	create_Time, _ := time.ParseInLocation("2006-01-02 15:04:05", createTime, time.UTC)
	nowTime := time.Now().UTC()
	sub := create_Time.Sub(nowTime)
	fmt.Println(time.Now().UTC())
	fmt.Println(int(sub.Hours()))
	opp := make([]int, 0)
	for i := 0; i < 10; i++ {
		opp = append(opp, len(opp))
	}
	fmt.Println(opp)
}
