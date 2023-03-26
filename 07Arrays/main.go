package main

import "fmt"

func main() {

	fmt.Println("Learning Arrays::")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tamatar"
	fruitList[3] = "Kela"

	fmt.Println("Konse fruit h thakur::", fruitList)
	fmt.Println("Kitene fruit h thakur::", len(fruitList))

	var sabjiList = [3]string{"Aloo", "Floweer", "Piyaz"}
	fmt.Println("Konse sabjiList h thakur::", sabjiList)
	fmt.Println("Kitene sabjiList h thakur::", len(sabjiList))
}
