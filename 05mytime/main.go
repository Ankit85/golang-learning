package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Learning Time in GO")

	currentTimeNow := time.Now()
	fmt.Printf("CurrentTimeNow: %v \n", currentTimeNow)

	createDate := time.Date(2025, time.January, 31, 15, 15, 15, 15, time.UTC)
	fmt.Printf("Creaated date:: %v \n", createDate)
	formatedDate(createDate)
}

func formatedDate(currentTimeNow time.Time) {
	formattedTime := currentTimeNow.Format("01-02-2006 Monday")
	fmt.Printf("formattedTime: %v \n", formattedTime)
}
	