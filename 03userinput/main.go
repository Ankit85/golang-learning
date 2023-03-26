package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome scam")

	fmt.Println("Enter the amount in dollar u want to scam::")

	input, _ := reader.ReadString('\n')

	fmt.Println("you like terminal is :", input)

}
