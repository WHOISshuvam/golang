package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://whoisbinit.me"

func main() {
	fmt.Println("Web Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T", response)
	defer response.Body.Close() // should close connection :)

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
