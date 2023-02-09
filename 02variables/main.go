package main

import "fmt"

const LoginToken string = "fjsdkljgkldjfkj" //Public

func main() {
	var username string = "suvam"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFLoat float64 = 255.3232357867866
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type: %T \n", smallFLoat)

	// default values and alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit type
	var website = "suvam.com"
	fmt.Println(website)

	//no var style
	numberofUsers := 30000000000.45
	fmt.Println(numberofUsers)
	fmt.Println(LoginToken)
	fmt.Printf("Variable type is of type: %T \n", LoginToken)
}
