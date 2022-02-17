package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	strat := time.Now()
	// 首先打开一个文件
	file, err := os.OpenFile("./resource.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// 然后读取文件里面的内容
	reader := bufio.NewReader(file)
	// 存入到缓冲区
	var str string
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str += readString
	}
	// 打开另外一个文件
	openFile, err := os.OpenFile("./copyTest.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	// 将缓冲区的内容写入到这个文件中
	writer := bufio.NewWriter(openFile)
	writer.WriteString(str)
	writer.Flush()
	end := time.Now()
	sub := end.Sub(strat).Nanoseconds()
	fmt.Println("一共消耗了", sub)
}
