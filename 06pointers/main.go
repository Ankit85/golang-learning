package main

import "fmt"

func main() {

	fmt.Println("Learning pointers")

	myNumber := 25

	var prts = &myNumber

	fmt.Println("Value of pointer is ", prts)
	fmt.Println("Acctual Value of pointer is ", *prts)

	*prts = *prts + 5
	fmt.Println("Now actual value is updated ::", myNumber)

}
