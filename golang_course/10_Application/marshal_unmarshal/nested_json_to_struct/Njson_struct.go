package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address struct {
		City string `json:"city"`
	} `json:"address"`
}

func main() {

	var userInfo UserInfo
	jsonStr := `{
    "ID": "1",
    "Name": "bilal shaikh",
    "Age": 30,
    "address": {
        "city": "pune"
    }
}`
	err1 := json.Unmarshal([]byte(jsonStr), &userInfo)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("% +v\n", &userInfo)

}
