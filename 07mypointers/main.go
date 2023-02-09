package main

import "fmt"

func main() {
	fmt.Println("Pointers Intro")
	// var ptr *int

	// fmt.Println("Value of Printer is", ptr)

	myNumber := 23
	// Reference &
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is: ", myNumber)
}
