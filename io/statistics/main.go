package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int // 英文的个数
	NumCount   int // 数字的个数
	SpaceCount int // 空格的个数
	OtherCount int // 其他字符的个数
}

func main() {
	// 如何实现统计一个文本里面的每个字符出现多少次
	// 定义一个结构体，用于保存统计的结果

	// 首先打开文件
	filePath := "./abc.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 定义一个CharCount的实例
	var count CharCount
	// 创建一个bufio的reader
	reader := bufio.NewReader(file)
	// 循环的读取文件的内容
	for true {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// 遍历str
		for _, val := range str {
			switch {
			case val >= 'a' && val <= 'z':
				fallthrough
			case val >= 'A' && val <= 'Z':
				count.ChCount++
			case val == ' ' || val == '\t':
				count.SpaceCount++
			case val >= '0' && val <= '9':
				count.NumCount++
			default:
				count.OtherCount++
				fmt.Println("其他字符是", string(val))
			}
		}
	}
	fmt.Println("字符的个数", count.ChCount, "数字的个数", count.NumCount, "空格的个数",
		count.SpaceCount, "其他字符", count.OtherCount)
}
