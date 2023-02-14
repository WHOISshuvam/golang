package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Functions in golang")

	result := adder(3, 5)

	proResult := proAdder(2, 3, 4, 5, 6)
	fmt.Println("Pro result is:", proResult)

	fmt.Println("Result is:", result)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Happy valentine's day")
}
