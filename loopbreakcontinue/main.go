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

	// for i :=0 0; i < 5; i++{
	// 	fmt.Println(i)
	// }

	name := []string{"Aaa", "Bbb", "Ccc", "Ddd", "fff"}

	for index, value := range name {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v\n", index, value)
	}

}
