package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Good Bye %v\n", n)
}

func planet(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Aryan")
	sayBye("Aryan")
	planet([]string{"pluto", "neptune", "mars"}, sayGreeting)

	a1 := circleArea(10.5)
	fmt.Println(a1)
	fmt.Printf("Circle 1  area is %0.3f", a1)
}
