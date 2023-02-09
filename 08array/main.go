package main

import "fmt"

func main() {

	fmt.Println("Arrays in golang")

	var friendsist [6]string
	friendsist[0] = "Arpan"
	friendsist[1] = "Ayush"
	friendsist[2] = "Anish"

	fmt.Println("Friends list is :", friendsist)
	fmt.Println("Friends list is :", len(friendsist)) // Out 6

	var cartoonsList = [3]string{"tom and jerry", "horrid henry", "ben 10"}
	fmt.Println("Cartoons list is:", cartoonsList)
	fmt.Println("Cartoons list is:", len(cartoonsList)) // out 3
}
