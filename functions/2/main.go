package main

import "fmt"

func add(x, y, z int) (int, int) {
	return x + y + z, x - y - z
}

func main() {

	ans1, ans2 := add(6, 78, 6)
	fmt.Println(ans1, ans2)
}
