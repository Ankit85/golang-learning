package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learning Slice")

	var mySliceList = []string{"Seb", "Amrud", "Santra"}

	fmt.Println("mySliceList", mySliceList)
	fmt.Printf("type of %T mySliceList  \n", mySliceList)

	mySliceList = append(mySliceList[1:3])
	fmt.Println("mySliceList", mySliceList)

	marks := make([]int, 4)
	marks[0] = 44
	marks[1] = 54
	marks[2] = 64
	marks[3] = 74

	fmt.Println("marks", marks)
	marks = append(marks, 11, 85, 999, 1425)
	fmt.Println("marks", marks)
	fmt.Println("marks", sort.IntsAreSorted(marks))
	sort.Ints(marks)
	fmt.Println("marks", marks)

	var courses = []string{"Java", "dart", "brainfuck", "javascript", "Typescript"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("updatted :", courses)

}
