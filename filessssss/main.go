package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling in golang")
	content := "This is a test text"

	file, err := os.Create("./testfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is ", length)
	file.Close()
}
