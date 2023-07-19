package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjpi"

func main() {
	fmt.Println("Welcome to Handling urls in Go")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	//fmt.Println(result.Scheme)
	//fmt.Println(result.Host)
	//fmt.Println(result.Path)
	//fmt.Println(result.Port())
	//fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The type of query params are: %T\n , qparams")

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsofURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		RawPath: "user=hitesh",
	}

	anotherURL := partsofURL.String()
	fmt.Println(anotherURL)

}
