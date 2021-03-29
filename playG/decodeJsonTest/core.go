package main

import (
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	fmt.Println("vim-go")
	a := `{"dsp101":{"u12345":{"age":19,"gender":1,"pos":"BJ"}}}`
	data := make(map[string]interface{})
	jsoniter.Unmarshal([]byte(a), &data)
	s := generateMappingResult(data, 100, false, false)
	fmt.Println(s)
}

func generateMappingResult(data interface{}, PackId int64, isPhoneNum, isMd5 bool) string {
	if isPhoneNum {
		if v, y := data.(string); y {
			return (v + "\t" + strconv.FormatInt(PackId, 10) + "\n")
		}
	} else {
		if v, y := data.(map[string]interface{}); y {
			for _, vv := range v {
				if vv == nil {
					return ""
				}
				if vvv, yyy := vv.(map[string]interface{}); yyy {
					for uid := range vvv {
						return (uid + "\t" + strconv.FormatInt(PackId, 10) + "\n")
					}
				}
			}
		}
	}
	return ""
}
