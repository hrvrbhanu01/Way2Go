package main

import "fmt"

func main() {
	var username string = "Bhanu"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4536748984899999
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)

}
