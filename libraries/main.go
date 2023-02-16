package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hellow world"
	fmt.Println(strings.Contains(greeting, "Hellow"))              // true
	fmt.Println(strings.ReplaceAll(greeting, "Hellow", "Destroy")) // Out : Destroy World
	fmt.Println(strings.ToUpper(greeting))                         // HELLOW WORLD
	fmt.Println(strings.Index(greeting, "ll"))                     // finds the occurence index of string
	fmt.Println(strings.Split(greeting, " "))                      // Converts string into array

	// Sort Package

	ages := []int{19, 21, 45, 65, 70, 34}
	sort.Ints(ages)
	fmt.Println(ages) //[19 21 34 45 65 70]

	index := sort.SearchInts(ages, 65)
	fmt.Println(index) /// search in sorted position

	names := []string{"Zpple", "Peach"}
	sort.Strings(names)
	fmt.Println(names) //[Peach Zpple]

}
