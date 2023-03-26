package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to name rating show!!")
	fmt.Println("How much do you rate a person whose name is Annie")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	cnumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic(err)
	}
	number := cnumber + 1
	fmt.Println(number)

}
