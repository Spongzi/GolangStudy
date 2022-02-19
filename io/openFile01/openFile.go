package main

import (
	"bufio"
	"os"
)

// 创建并写入操作

func main() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	str := "Hello\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		// 将内容先写到缓存区
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}
	// 读取缓存区的内容
	err = writer.Flush()
	if err != nil {
		return
	}
}
