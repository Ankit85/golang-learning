package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Learning files")

	content := "this iis the content of file i m putting"

	file, err := os.Create("./myFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length of fileis:", length)

	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("File has content \n", dataByte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
