package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Rose day", "Kiss Day", "Propose Day", "Promise Day", "Teddy Day"}

	// fmt.Println(days)
	// for g := 0; g < len(days); g++ {
	// 	fmt.Println(days[g])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	suvam := 1
	for suvam < 10 {
		fmt.Println("Value is", suvam)
		suvam++
	}

}
