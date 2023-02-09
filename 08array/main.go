package main

import "fmt"

func main() {

	fmt.Println("Arrays in golang")

	var friendsList [6]string

	friendsList[0] = "Reason"
	friendsList[1] = "Rohan"
	friendsList[3] = "Ritesh"

	fmt.Println("Friends List is :", friendsList)
	fmt.Println("Friends List is :", len(friendsList)) // output 6

	var cartoonsList = [3]string{"Ben10", "Horrid Henry", "Tom and Jerry"}

	fmt.Println("Cartoons list is", len(cartoonsList))
}
