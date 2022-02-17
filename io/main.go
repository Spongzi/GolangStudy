package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../README.md")
	if err != nil {
		fmt.Println("Open File failed", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Close file failed", err)
		}
	}(file)
	fmt.Printf("file = %v", file)
}
