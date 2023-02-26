package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {

	a, err := os.Open("abc.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	b, err := ioutil.ReadAll(a)
	if err != nil {
		log.Fatal(err)
	}

	abc := string(b)

	newQR := "abcd-qr.png"
	err = qrcode.WriteFile(abc, qrcode.Medium, 512, newQR)
	if err != nil {
		log.Fatal(err)
	}
}
