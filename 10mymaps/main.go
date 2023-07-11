package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["GO"] = "Golang"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "PY")
	fmt.Println("List of all languages:", languages)

	//loops in golang

	for key, value := range languages {
		fmt.Println("For key %v, value is %v\n", key, value)
	}

}
