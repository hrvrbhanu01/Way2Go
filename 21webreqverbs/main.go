package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to we request verbs!")
}

func PerformGetRequest() {
	const myurl = ""

	response, err := http.Get(myurl)

}
