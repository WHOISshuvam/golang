package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	defer file.Close()
	readFile("./testfile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data in the file is \n", string(dataByte))
}
