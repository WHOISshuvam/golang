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

	if age >= 10 {
		fmt.Println("You are eligible to play")
	} else {
		fmt.Println("You can't play")
		return
	}

	score := 0

	fmt.Printf("How do you see hidden character in text file in linux? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Println(answer, answer2)

	if answer+" "+answer2 == "cat -v" {
		fmt.Println("Correct")

	} else if answer+" "+answer2 == "Cat -v" {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("What is the capital city of Nepal? ")
	var capital string
	fmt.Scan(&capital)

	if answer == "kathmandu" {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}
}
