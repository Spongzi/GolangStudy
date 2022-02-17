package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 读取整个文件， 但是不适合读取一个很大的文件
func main() {
	readFile, err := os.ReadFile("../README.md")
	if err != nil {
		return
	}
	file, err := ioutil.ReadFile("../README.md")
	if err != nil {
		fmt.Println("ReadFile failed!", err)
		return
	}
	fmt.Println(string(file))
	fmt.Println("================")
	fmt.Println(string(readFile))
}
