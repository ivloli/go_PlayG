package main

import (
	"encoding/json"
	"fmt"
)

type Org struct {
	Id       string
	Name     string
	Children *[]*Org `json:"children,omitempty"`
}
type Item struct {
	Id    string
	Name  string
	Pid   string
	Level int
}

func main() {
	OmapSlice := make([]map[string]*Org, 0)
	rootMap := make(map[string]*Org)
	rootMap["1"] = &Org{
		Id:   "1",
		Name: "Operator",
	}
	children := make([]*Org, 0)
	rootMap["1"].Children = &children
	OmapSlice = append(OmapSlice, rootMap)
	L2 := []Item{
		Item{"o20000001", "suGroup", "1", 1},
		Item{"o20000002", "wiGroup", "1", 1},
		Item{"o20000003", "sqGroup", "1", 1},
	}
	setOrg(&OmapSlice, L2)
	L3 := []Item{
		Item{"o30000001", "webTeam", "o20000001", 2},
		Item{"o30000002", "UiTeam", "o20000001", 2},
		Item{"o30000003", "webTeam", "o20000002", 2},
	}
	setOrg(&OmapSlice, L3)
	L4 := []Item{
		Item{"o40000001", "webTeam_eric", "o30000001", 3},
		Item{"o40000002", "webTeam_adam", "o30000001", 3},
		Item{"o40000003", "UiTeam_whu", "o30000002", 3},
		Item{"o40000004", "webTeam_lee", "o30000003", 3},
	}
	setOrg(&OmapSlice, L4)
	L5 := []Item{
		Item{"o50000001", "eric_lu_team", "o40000001", 4},
	}
	setOrg(&OmapSlice, L5)
	sssss, _ := json.Marshal(rootMap["1"])
	fmt.Println(string(sssss))
	var empty []Org
	for _, v := range empty {
		fmt.Println("??")
		fmt.Println(v)
	}
}

func setOrg(mapSlice *[]map[string]*Org, L []Item) bool {
	levelMap := make(map[string]*Org)
	for _, item := range L {
		newOrg := &Org{
			Id:   item.Id,
			Name: item.Name,
		}
		rootMap := (*mapSlice)[item.Level-1]
		if rootMap == nil || len(rootMap) == 0 {
			return false
		}
		if val, ok := rootMap[item.Pid]; ok {
			if val.Children == nil {
				children := make([]*Org, 0)
				val.Children = &children
			}
			*val.Children = append(*val.Children, newOrg)
		} else {
			return false
		}
		levelMap[item.Id] = newOrg
	}
	*mapSlice = append(*mapSlice, levelMap)
	return true
}
