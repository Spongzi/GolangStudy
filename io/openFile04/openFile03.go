package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for true {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(readString)
	}
	write := bufio.NewWriter(file)
	str := "你好北京\n"
	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	write.Flush()
}
