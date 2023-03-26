package main

import (
	"fmt"
	"net/http"

	"github.com/ankit85/movieapi/router"
)

func main() {

	fmt.Println("Here we go again")

	r := router.Router()

	http.ListenAndServe(":3000", r)

}
