package main

import (
	"fmt"
	"strings"
	"unicode"
)

type ContactInfos struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

func main() {
	fmt.Println("vim-go")
	a := []string{"hailong:1233123212312;xia:1111", "hailong1:12313123212312;s:1111+"}
	b := make([]ContactInfos, 0)
	for _, item := range a {
		b = append(b, getContactInfos(item)...)
	}
	fmt.Println(checkDuplicatePhoneNumber(b))
}

func checkDuplicatePhoneNumber(infos []ContactInfos) bool {
	fmt.Println(infos)
	tmap := make(map[string]struct{}, len(infos))
	for _, item := range infos {
		if _, ok := tmap[GetStringOnlyIncludeNumber(item.Number)]; ok {
			return false
		} else {
			tmap[GetStringOnlyIncludeNumber(item.Number)] = struct{}{}
		}
	}
	return true
}
func GetStringOnlyIncludeNumber(s string) string {
	res := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsDigit(r) {
			res = append(res, r)
		}
	}
	return string(res)
}
func getContactInfos(contactStr string) []ContactInfos {
	var records []ContactInfos
	records = make([]ContactInfos, 0)
	segs := strings.Split(contactStr, ";")
	for _, item := range segs {
		subsegs := strings.Split(item, ":")
		if len(subsegs) == 2 {
			record := ContactInfos{
				Name:   subsegs[0],
				Number: subsegs[1],
			}
			records = append(records, record)
		}
	}
	return records
}
