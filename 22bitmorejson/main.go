package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` //this will hide the password
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"Go Programming", 10, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"Go", 13, "Udemy", "abcd123", []string{"dev", "js"}},
		{"Solidity Programming", 19, "Udemy", "abcrt123", []string{"blockchain-dev", "Go"}},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //there is Marshal as well
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Solidity Programming",
		"Price": 19,
		"website": "Udemy",
		"tags": ["blockchain-dev","Go"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else {
		fmt.Println("Invalid Json Data")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v:= range myOnlineData{
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
