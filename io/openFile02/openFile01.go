package main

import (
	bufio "bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Open file failed", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "你好苏旭\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
