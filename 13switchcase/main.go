package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in Go")

	rand.Seed(time.Now().UnixNano()) //for random number evrytime we roll the dice.
	diceNumber := rand.Intn(6) + 1   //we want number between 1 to 6 but the range is non-inclusive so we need to add 1 to include 6 as well.
	fmt.Println("Value of Dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough //when we get a 3 it will also print the case 4 with it.
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and play again")
	default: //to handle all the rest of the cases which could be there!
		fmt.Println("What was that!?")

	}

}
