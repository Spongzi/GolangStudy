package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// TODO 写一个函数 接受文件路径
func copyFile(src string, dst string) (write int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("open file err", err)
		return 0, err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return 0, err
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}
func main() {
	srcFile := "../test.txt"
	dstFile := "./copy"
	_, err := copyFile(srcFile, dstFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("err", err)
	}
}
