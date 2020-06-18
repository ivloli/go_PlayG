package main

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
)

type sss struct {
	Ulrs []string `json:"ulrs"`
	SB   []string `json:"sb"`
}

type ContactInfos struct {
	Name   string `json:"name"`
	Number string `json:"number"`
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

func main() {
	adv := "sme123469090"
	fmt.Println(adv[0:3])
	fmt.Println(adv[0:3] == "sme")
	str := `{"ulrs":["aaaaa","bbbbbb"]}`
	var s sss
	err := json.Unmarshal([]byte(str), &s)
	fmt.Println(err)
	s.SB = append(s.SB, "sb001")
	fmt.Println(s)
	ppp := "/api/uploaded_file/get/pic/2020/2/18/2540589857792.jpg"
	fmt.Printf(path.Join("/nfs/am", ppp[22:len(ppp)]))
	fmt.Printf(ppp[23:len(ppp)])
	tmap := make(map[string]string, 7)
	tmap["xiaogou"] = "dog"
	fmt.Println("--------------")
	if tmap["xiaoji"] == "" {
		fmt.Println("none")
	}
	fmt.Println(getContactInfos("12312:123123123"))
}
