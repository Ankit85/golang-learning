package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Learning swtich")

	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("You can open the board number is", diceNumber)
	case 2:
		fmt.Println("You can move the board number is", diceNumber)
	case 3:
		fmt.Println("You can move the board number is", diceNumber)
	case 4:
		fmt.Println("You can move the board number is", diceNumber)
		fallthrough
	case 5:
		fmt.Println("You can move the board number is", diceNumber)
	case 6:
		fmt.Println("You can move and roll again as the board number is", diceNumber)
	default:
		fmt.Println("What the fuck that was!!!")

	}

}
