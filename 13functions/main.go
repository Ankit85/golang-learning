package main

import "fmt"

func main() {

	fmt.Println("Learning function")

	greetings()
	addN := addTwoNum(444444, 888888)
	fmt.Println(addN)
	mulN, msg := mulNum(4, 4)
	fmt.Println(mulN)
	fmt.Println(msg)
}

func greetings() {
	fmt.Println("Welcome to Costco!!")
}

func addTwoNum(num1 int, num2 int) int {
	return num1 + num2
}

func mulNum(values ...int) (int, string) {

	total := 1

	for _, value := range values {
		total = total * value
	}
	return total, "I did it"
}
