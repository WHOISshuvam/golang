package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	password string   `json:"-"`
	Tags     []string `json:"tags, omitempty`
}

func main() {
	fmt.Println("Json")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs BootCamp", 299, "Learn Code Online", "Password123", []string{"web-dev", "js"}},
		{"MERN BootCamp", 299, "Learn Code Online", "Password1234", []string{"full-stack", "js"}},
		{"Angular BootCamp", 299, "Learn Code Online", "Password1235", []string{"google", "js"}},
	}

	//package this data as Json Data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
