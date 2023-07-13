package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two") //last in first out
	fmt.Println("Hello")
	myDefer()
}

// world,one, two,hello <<<

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//Hello will be printed then myDefer() is in the line then deferFmt will be printed as last in first out
