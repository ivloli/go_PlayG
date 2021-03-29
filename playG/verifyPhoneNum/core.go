package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(VerifyMobileFormat("13811995547"))
	fmt.Println(VerifyMobileFormat("19411995547"))
	fmt.Println(VerifyMobileFormat("1851195547"))
}

func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7,9])|(15[0-3,5-9])|(16[2,5,7,9])|(17[0,3,5-8])|(18[0-9])|(19[0-3,5-9]))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
