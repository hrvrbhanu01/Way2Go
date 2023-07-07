package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok syntax|| err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating", input)
	fmt.Printf("type of the rating is %T", input)

}
