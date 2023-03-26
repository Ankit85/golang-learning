package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("my")
	defer fmt.Println("name")
	defer fmt.Println("is ")
	defer fmt.Println("khan")
	myCustomDefer()
}
func myCustomDefer() {
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
}
