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
	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"Go Programming", 10, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"Go", 13, "Udemy", "abcd123", []string{"dev", "js"}},
		{"Solidity Programming", 19, "Udemy", "abcrt123", []string{"blockchain-dev", "Go"}},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
