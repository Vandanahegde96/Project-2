package main

import (
	"fmt"
	"sort"
	"encoding/json"
)

func main() {
	type Emp struct {
		Id  string `json:"Id"`
		Name string `json:"Name"`
	}

	var emps []Emp

	users := `[{"Id":"1001","Name":"Kerry Morton"},{"Id":"1002","Name":"Erica Bowers"},{"Id":"1003","Name":"Conrad Patton"}]`
	err :=json.Unmarshal([]byte(users),&emps)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(emps, func(i,j int)bool{
		return emps[i].Name < emps[j].Name
	})

	fmt.Println(emps)
}