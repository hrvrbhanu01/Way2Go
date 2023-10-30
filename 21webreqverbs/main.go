package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request video!")
	//PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(content)
	// fmt.Println(string(content))

}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"name": "John",
		"age": 30,
		"city": "New York"
	}`)

	response,err:= http.Post(myurl, "application/json", requestBody)

	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()

	content, _:= ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
