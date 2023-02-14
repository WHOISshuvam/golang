package main

import "fmt"

func main() {
	defer fmt.Println("fffffffffkkkkkkkkkkkkkkkkkk")
	defer fmt.Println("fffffffffkkkkkkkkkkkkkkkkkk1")
	defer fmt.Println("fffffffffkkkkkkkkkkkkkkkkkk2")
	defer fmt.Println("fffffffffkkkkkkkkkkkkkkkkkk3")
	fmt.Println("Hello World")
	myDefer()

	//output hello world 43210 /n fffffffffkkkkkkkkkk 3 2 1 lifoooo
}

func myDefer() {
	for i := 0; i < 5; i++ {
		// stack 012345 >> difo 43210
		defer fmt.Print(i)
	}
}
