package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 构造缓冲区，可以读取一个很大的文件

func main() {
	file, err := os.Open("../README.md")
	if err != nil {
		fmt.Println("open file failed!", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close file failed!", err)
		}
	}(file)
	reader := bufio.NewReader(file)
	// 循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF代表文件的末尾
			return
		}
		fmt.Println(str)
	}
}
