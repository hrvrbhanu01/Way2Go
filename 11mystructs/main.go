package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	//no inheritence in go; No super or parent or no child i.e.classes are not there in go.

	hitesh := User{"Bhanu", "hitesh@go.dev", true, 18}
	fmt.Println(hitesh)
	fmt.Println("hitesh details are: %+v\n", hitesh)
	fmt.Println("Name is %v and email is %v.", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
