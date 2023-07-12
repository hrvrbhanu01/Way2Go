package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in Go")
	greeter()

	Result := adder(3, 5)
	fmt.Println("Result is: ", Result)

	proRes, myMessage := proAdder(2, 3, 5, 6)
	fmt.Println("Pro Result is : ", proRes)
	fmt.Println("My message is :", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi pro result here"
}

func greeter() {
	fmt.Println("Namaste from Golang")
}
