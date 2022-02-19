package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	str := "abc"
	_, err = writer.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush failed!", err)
		return
	}
}
