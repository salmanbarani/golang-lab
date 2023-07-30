package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[len(os.Args)-1] // slice of type string
	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("There was an error: ", error.Error())
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
