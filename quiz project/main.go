package main

import "fmt"

func main() {

	fmt.Println("Hello Player! Best of Luck :) ")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v, Welcome to the game \n", name)

	fmt.Printf("Enter your age: \n")
	var age uint
	fmt.Scan(&age)

	fmt.Println(age >= 10)
}
