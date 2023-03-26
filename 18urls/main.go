package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=cgnhhkjasda"

func main() {

	fmt.Println("Learning urls")

	result, _ := url.Parse(myurl)
	/* fmt.Println("resule", result.Scheme)
	fmt.Println("resule", result.Host)
	fmt.Println("resule", result.Path)
	fmt.Println("resule", result.RawQuery)
	*/
	query := result.Query()
	fmt.Println("query", query["payment"])
	fmt.Println("query", query["coursename"])
	for _, param := range query {
		fmt.Println(param)
	}

	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
