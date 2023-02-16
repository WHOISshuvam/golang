package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in golang")
	var friendsList = []string{"Bishal", "Bikal", "Binit"}
	fmt.Printf("The type of friendsList is %T\n", friendsList)

	friendsList = append(friendsList, "Bishaka", "Binita")
	fmt.Println(friendsList)

	friendsList = append(friendsList[:3])
	fmt.Println(friendsList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 934
	highScore[2] = 534
	highScore[3] = 834
	//fmt.Println(highScore) >> Crashes program
	highScore = append(highScore, 123, 456, 789)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore)) // true

	var scores = []int{100, 20}
	scores[1] = 25
	scores = append(scores, 78)

	fmt.Println(scores, len(scores))
}
