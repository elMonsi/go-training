package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// filename := os.Args[1]
	// data, err := os.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println("Not able to read file.", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(data))

	//using interfaces
	filename := os.Args[1]
	data, err := os.Open(filename)
	if err != nil {
		fmt.Println("Not able to read file.", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, data)
}
