package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Go")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	for _, day := range days {
		fmt.Printf("index is and value is %v\n", day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 4 {
			goto lco
		}

		if rougueValue == 5 {
			break
		}

		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at LearnOnline")
}
