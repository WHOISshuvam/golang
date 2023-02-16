package main

import "fmt"

func main() {

	age := 21
	name := "oooooooooooooooooooo"
	marks := 97.3432

	fmt.Printf("My name is %v and my age is %v\n", name, age)
	fmt.Printf("My name is %q and my age is %q\n", name, age)
	fmt.Printf("Score is %0.2f", marks) // Rounds off decimal to 2 place

	// Sprintf (Save  formatted strings)
	var str = fmt.Sprintf("My name is %q and my age is %q\n", name, age)
	fmt.Println("The saved string is", str)
}
