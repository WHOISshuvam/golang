package main

import (
	"fmt"
	"net/url"
)

func main() {

	const myUrl string = "https://whoisbinit.me:443?_next=another.comwebsite&year=2023"

	fmt.Println("URL handling in golang")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Printf("The type of query param are %T\n", qparam)

	fmt.Println(qparam["year"])

	for _, value := range qparam {
		fmt.Println("Param is ", value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "whoisbinit.me",
		Path:   "null",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
