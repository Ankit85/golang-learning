package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	dataByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	fmt.Println(content)

}
