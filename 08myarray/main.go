package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Go")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[2] = "mango"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie List is: ", vegList)
	fmt.Println("Veggie List is: ", len(vegList))

}
