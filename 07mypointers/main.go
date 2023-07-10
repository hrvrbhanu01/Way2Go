package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	//var ptr *int

	//fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber //referencing a pointer to a number that's why using "&"

	fmt.Println("Value of actual pointer is ", ptr)  //refers to the address to the pointer
	fmt.Println("Value of actual pointer is ", *ptr) //refers to what value inside the pointer

	*ptr = *ptr * 2
	fmt.Println("New Value is: ", myNumber)

}
