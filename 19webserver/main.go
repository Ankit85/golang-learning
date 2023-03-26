package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const murl string = "http://localhost:3000"

func main() {

	fmt.Println("Learning webrequest methods")

	// PerformGetMethods(url)
	// PerformPostMethod(url)
	PerformPostMethodUsingFormData(murl)

}

func PerformGetMethods(url string) {
	response, err := http.Get(url + "/get")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var responseBuilder strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteData, _ := responseBuilder.Write(content)
	if response.StatusCode == 200 {
		fmt.Println("byteData", byteData)
		fmt.Println("response string ", responseBuilder.String())
	} else if response.StatusCode == 404 {
		fmt.Println("404 Not found")
	}
}

func PerformPostMethod(url string) {

	postPayload := strings.NewReader(`
	{
		"cname":"Go Lang",
		"price": 0,
		"platform":"youtube" 
	}
	`)

	response, err := http.Post(url+"/post", "application/json", postPayload)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("response :", string(content))
}

func PerformPostMethodUsingFormData(wurl string) {

	postFormData := url.Values{}
	postFormData.Add("FName", "Ankitwa")
	postFormData.Add("LName", "Viswaa")
	postFormData.Add("Loc", "India")

	response, err := http.PostForm(wurl+"/postform", postFormData)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println("response :", string(content))
}
