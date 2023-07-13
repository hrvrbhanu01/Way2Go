package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	//no inheritence in go; No super or parent or no child i.e.classes are not there in go.

	hitesh := User{"Bhanu", "hitesh@go.dev", true, 18}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email %v\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is : ", u.Email)
}
