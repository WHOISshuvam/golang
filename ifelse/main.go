package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	requests := 10000
	var result string
	if requests < 10000 {
		result = "Every thing is normal"
	} else if requests > 10000 {
		result = "Security Alert"
	} else {
		result = "Acceptable"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
