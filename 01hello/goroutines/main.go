package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointers

func main() {
	// go greeter("Hello") // Go routine
	// greeter("World")

	websitelist := []string{"https://go.dev",
		"https://evil.com",
		"https://hackerone.com",
		"https://tesla.com",
		"https://nonexistingsiteooo.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2 * time.Millisecond)
// 		fmt.Println(s)
// 	}
//}

func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("\n")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
