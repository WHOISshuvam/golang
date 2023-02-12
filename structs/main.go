package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	// no inheritace in golang; No super or parents
	suvam := User{"suvam", "suvam@trendmicro.com", true, 19}
	fmt.Println(suvam)
	fmt.Printf("suvam details are %+v\n", suvam)
	fmt.Printf("Name is %v and email is %v\n", suvam.Name, suvam.Email)

}

type User struct {
	Name    string
	Email   string
	isAdmin bool
	Age     int
}
