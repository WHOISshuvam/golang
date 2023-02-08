package main

import "fmt"

var jwtToken uint64 = 1200

const LoginToken string = "gjhgdshf"

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.5385787598375
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	// default values and alias

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of  type: %T \n", anotherVariable)

	//implicit

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberofUsers := 300000
	fmt.Println(numberofUsers)

	fmt.Println(LoginTokens)
	fmt.Printf("Variable is of type %T \n", LoginTokens)
}
