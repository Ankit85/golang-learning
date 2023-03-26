package main

import (
	"fmt"
)

func main() {

	fmt.Println("Learning maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["BF"] = "BrainFcuk"
	languages["TY"] = "Typescript"

	// aacessing the map list
	fmt.Println("languages::", languages)

	// accessing values from Key
	fmt.Println("Accessing value using key:", languages["JS"])

	// deleting values
	delete(languages, "BF")
	fmt.Println("languages::", languages)

	// using for loop
	for _, value := range languages {
		fmt.Printf("Language has value %v \n", value)
	}

}
